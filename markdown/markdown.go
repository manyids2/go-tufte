package markdown

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
)

// Datastructure to hold one markdown document
type Document struct {
	// Path to markdown file
	Path string

	// Data of file ( we read the whole file )
	Buffer *[]byte

	// Level 1 header, warn, set to path if not found
	Title string

	// Tree containing info about sections, paras, etc.
	Root *Section

	// Sections as list, with levels
	Sections []*Section
}

// NewDocument Reads new markdown file and tries to parse
// Returns error on failure
func NewDocument(path string) (*Document, error) {
	doc := Document{
		Path: path,
	}

	// Read file, return if invalid
	content, err := os.ReadFile(doc.Path)
	if err != nil {
		return nil, err
	}
	doc.Buffer = &content

	// Parse source code
	parser := sitter.NewParser()
	parser.SetLanguage(markdown.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, content)
	if err != nil {
		return nil, err
	}
	n := tree.RootNode()

	// TODO: approximating last row by last child of root
	// fails if root has no children ( like empty file )
	doc.Root = &Section{
		Node:      n,
		Level:     0,
		StartByte: 0,
		StartRow:  0,
		EndByte:   uint32(len(content)),
		EndRow:    0,
	}

	// Return if new file
	if n.ChildCount() == 0 {
		return &doc, nil
	}

	// Else compute last row ( assumed finished before next step )
	doc.Root.EndRow = n.Child(int(n.ChildCount()) - 1).EndPoint().Row

	// Set sections
	doc.SetSections()

	// Set Root with proper children
	doc.SetRootFromSections()

	// Set title to first found header
	if len(doc.Root.Children) > 0 {
		doc.Title = doc.Root.Children[0].Title
	}

	return &doc, nil
}

// Tree of sections
type Section struct {
	// Reference to tree-sitter node of atx_heading
	Node *sitter.Node

	// Heading level
	Level int

	// Heading text
	Title string

	// Byte boundaries for reading/writing
	StartByte, EndByte uint32

	// Row boundaries for display ('\n' is parsed already)
	StartRow, EndRow uint32

	// Parent section
	Parent *Section

	// There may be content above children
	Content []*Content

	// Children sections.. can there be content after child?
	Children []*Section
}

func (d *Document) NewSection(n *sitter.Node) *Section {
	// Parse level
	level, err := strconv.Atoi(n.Child(0).Type()[5:6]) // Header type (1/2/3/... with hash)
	if err != nil {
		level = 0 // Floating
	}

	// Parse title
	title := n.Child(1).Child(0).Content(*d.Buffer)
	title = strings.Trim(title, " ") // Necessary for headers

	// Create new section
	return &Section{
		Node:      n,
		Level:     level,
		Title:     title,
		StartRow:  n.StartPoint().Row,
		StartByte: n.StartByte(),
	}
}

func (d *Document) SetSections() {
	n := d.Root.Node

	// Allocates new list each time it is called
	sections := make([]*Section, 0, n.ChildCount())

	// NOTE: Here is where parsing breaks down
	// We need to walk to find all the sections
	for i := 0; i < int(n.ChildCount()); i++ {
		if n.Child(i).Type() == "atx_heading" {
			section := d.NewSection(n.Child(i))
			sections = append(sections, section)
		}
	}

	if len(sections) < 1 {
		return
	}

	// Reiterate to set end rows
	for i := 0; i < len(sections)-1; i++ {
		sections[i].EndRow = sections[i+1].StartRow - 1
		sections[i].EndByte = sections[i+1].StartByte - 1
	}
	sections[len(sections)-1].EndRow = d.Root.EndRow
	sections[len(sections)-1].EndByte = d.Root.EndByte

	// Set on document
	d.Sections = sections
}

func (d *Document) GetMaxLevel() int {
	level := 0
	for _, s := range d.Sections {
		if s.Level > level {
			level = s.Level
		}
	}
	return level
}

// Modifies Root based on Sections - should be called inside set sections
func (d *Document) SetRootFromSections() {
	// Variables to keep track
	previous := d.Root
	maxLevel := d.GetMaxLevel()
	levelLatest := make([]*Section, maxLevel+1, maxLevel+1)
	levelLatest[0] = d.Root

	// Iterate over headers and put into tree
	for _, n := range d.Sections {
		if previous.Level == n.Level {
			n.Parent = previous.Parent
		} else if previous.Level < n.Level {
			n.Parent = previous
		} else if previous.Level > n.Level {
			parent := previous.Parent
			for {
				if parent.Level >= n.Level {
					parent = parent.Parent
				} else {
					break
				}
			}
			n.Parent = parent
		}
		n.Parent.Children = append(n.Parent.Children, n)
		levelLatest[n.Level] = n
		previous = n
	}
}

// WalkWithIndent sections
func (n *Section) WalkWithIndent(indent int, callback func(n *Section, indent int)) {
	for _, child := range n.Children {
		callback(child, indent)
		child.WalkWithIndent(indent+2, callback)
	}
}

// Print to stdout
func (d *Document) PrintSections(w io.Writer) {
	d.Root.WalkWithIndent(0, func(n *Section, indent int) {
		s := strings.Repeat(" ", indent)
		fmt.Fprintf(w, "%s%s ", s, n.Title)
		fmt.Fprintf(w,
			"|> %d [%d, %d] (%d, %d)\n",
			n.Level,
			n.StartRow, n.EndRow,
			n.StartByte, n.EndByte,
		)
	})
}

// Single block of content ( paragraph, code-block, blockquotes, etc. )
type Content struct {
	Node               *sitter.Node
	Index              int
	StartByte, EndByte uint32
	StartRow, EndRow   uint32
	Parent             *Section
}
