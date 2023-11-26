package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
	"github.com/spf13/cobra"
)

func ToC(path string) {
	// Read file
	dat, err := os.ReadFile(path)
	CheckErr(err)

	// Create parser
	parser := sitter.NewParser()
	parser.SetLanguage(markdown.GetLanguage())

	// Parse source code
	tree := parser.Parse(nil, dat)

	// Root node
	n := tree.RootNode()

	// Store nodes of headers
	headers := []*sitter.Node{}

	// TODO: Can make a struct here
	levels := []int{}
	startRows := []int{}
	startBytes := []int{}
	endRows := []int{}
	endBytes := []int{}

	// Walk from root and store tree-sitter nodes, boundaries
	WalkWithIndent(n, 0, func(n *sitter.Node, indent int) {
		if n.Type() == "atx_heading" {
			level, err := strconv.Atoi(n.Child(0).Type()[5:6]) // Header type (1/2/3/... with hash)
			CheckErr(err)
			headers = append(headers, n)
			levels = append(levels, level)
			startRows = append(startRows, int(n.StartPoint().Row))
			startBytes = append(startBytes, int(n.StartByte()))
		}
	})
	for i := 0; i < len(levels)-1; i++ {
		endRows = append(endRows, startRows[i+1]-1)
		endBytes = append(endBytes, startBytes[i+1]-1)
	}
	// TODO: approximating last row by last element
	lastRow := n.NamedChild(int(n.NamedChildCount()) - 1).EndPoint().Row
	// lastColumn := n.NamedChild(int(n.NamedChildCount()) - 1).EndPoint().Column
	endRows = append(endRows, int(lastRow))
	endBytes = append(endBytes, len(dat))

	// Ceate node tree
	var previousNodeTree *NodeTree
	maxLevel := MaxOverArray(levels)
	levelLatest := make([]*NodeTree, maxLevel+1, maxLevel+1)

	// Root
	t := &NodeTree{
		Level:    0,
		Parent:   nil,
		Children: []*NodeTree{},
	}
	levelLatest[0] = t
	previousNodeTree = t

	// Iterate over headers and put into tree
	for i, n := range headers {
		// NOTE: Copying = allocating
		text := n.Child(1).Child(0).Content(dat)
		text = strings.Trim(text, " ") // Necessary for headers
		nt := &NodeTree{
			Text:        text,
			StartByte:   startBytes[i],
			EndByte:     endBytes[i],
			StartRow:    startRows[i],
			EndRow:      endRows[i],
			StartColumn: 0,
			EndColumn:   0, // unknown
			Level:       levels[i],
			Children:    []*NodeTree{},
		}

		// Check is same level / lower / higher
		if previousNodeTree.Level == nt.Level {
			nt.Parent = previousNodeTree.Parent
		} else if previousNodeTree.Level < nt.Level {
			nt.Parent = previousNodeTree
		} else if previousNodeTree.Level > nt.Level {
			parent := previousNodeTree.Parent
			for {
				if parent.Level >= nt.Level {
					parent = parent.Parent
				} else {
					break
				}
			}
			nt.Parent = parent
		}
		nt.Parent.Children = append(nt.Parent.Children, nt)

		// Update tracking variables
		levelLatest[nt.Level] = nt
		previousNodeTree = nt
	}

	// Print nodes with text, indent to check
	WalkNodeTree(t, 0, CallbackPrintText)
}

// tocCmd represents the cli command
var tocCmd = &cobra.Command{
	Use:   "toc",
	Short: "ToC",
	Long:  `Table of contents`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Please specify --path.")
		}
		ToC(path)
	},
}

func init() {
	rootCmd.AddCommand(tocCmd)
}
