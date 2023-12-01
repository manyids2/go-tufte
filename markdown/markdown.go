package markdown

import (
	"context"
	"fmt"
	"io"
	"log"
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
	doc.Root = doc.NewSection(tree.RootNode())

	// Return if new file
	if doc.Root.Node.ChildCount() == 0 {
		return &doc, nil
	}

	// Set sections
	doc.SetSections()

	// Set Root with proper children
	doc.SetRootFromSections()

	// Set title to first found header
	if len(doc.Sections) > 0 {
		doc.Title = doc.Sections[0].Title
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

	// Boundaries for display ('\n' is parsed already)
	StartRow, StartColumn uint32
	EndRow, EndColumn     uint32

	// Parent section
	Parent *Section

	// There may be content above children
	Content []*Content

	// Children sections.. can there be content after child?
	Children []*Section
}

func (d *Document) NewSection(n *sitter.Node) *Section {
	// Parse level
	level, err := strconv.Atoi(n.NamedChild(0).NamedChild(0).Type()[5:6]) // Header type (1/2/3/... with hash)
	if err != nil {
		level = 0 // Floating
	}

	// Parse title
	title := n.NamedChild(0).NamedChild(1).Content(*d.Buffer)
	title = strings.Trim(title, " ") // Necessary for headers

	// Create new section
	return &Section{
		Node:        n,
		Level:       level,
		Title:       title,
		StartRow:    n.StartPoint().Row,
		EndRow:      n.EndPoint().Row,
		StartColumn: n.StartPoint().Column,
		EndColumn:   n.EndPoint().Column,
		StartByte:   n.StartByte(),
		EndByte:     n.EndByte(),
	}
}

func (d *Document) SetSections() {
	// Allocates new list each time it is called
	sections := []*Section{}

	// NOTE: Here is where parsing breaks down
	// We need to walk to find all the sections
	WalkTSNode(d.Root.Node, func(n *sitter.Node) {
		if n.Type() == "section" {
			section := d.NewSection(n)
			sections = append(sections, section)
		}
	})

	// Error if nothing found
	if len(sections) < 1 {
		log.Fatal("No sections found")
		return
	} else {
		fmt.Printf("%d sections found\n", len(sections))
	}

	d.SetRootFromSections()

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

// WalkTSNode Walk tree-sitter node
func WalkTSNode(n *sitter.Node, callback func(n *sitter.Node)) {
	for i := 0; i < int(n.NamedChildCount()); i++ {
		child := n.NamedChild(i)
		callback(child)
		WalkTSNode(child, callback)
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
