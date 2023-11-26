package cmd

import (
	"fmt"
	"os"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var path string

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Create parser
		parser := sitter.NewParser()
		parser.SetLanguage(markdown.GetLanguage())

		// Read file
		dat, err := os.ReadFile(path)
		check(err)

		// Parse source code
		tree := parser.Parse(nil, dat)

		// Query tree
		n := tree.RootNode()

		// Print in treesitter format
		// fmt.Println(n)

		// Get children ( not nested )
		for i := 0; i < int(n.NamedChildCount()); i++ {
			child := n.NamedChild(i)
			fmt.Println(child.Type(), child.StartByte(), child.EndByte())
			fmt.Println(child.Content(dat))
		}
	},
}

func init() {
	cliCmd.PersistentFlags().StringVarP(&path, "path", "p", "./README.md", "Path to markdown file.")
	rootCmd.AddCommand(cliCmd)
}
