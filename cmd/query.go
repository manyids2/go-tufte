package cmd

import (
	"fmt"
	"log"
	"os"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Please specify --path.")
		}

		// Create parser
		parser := sitter.NewParser()
		parser.SetLanguage(markdown.GetLanguage())

		// Read file
		dat, err := os.ReadFile(path)
		CheckErr(err)

		// Parse source code
		tree := parser.Parse(nil, dat)

		// Root node
		n := tree.RootNode()

		// Compile query
		lang := markdown.GetLanguage()
		q, err := sitter.NewQuery([]byte(
			`((setext_heading)
				(heading_content 
				  text:	(text) @the-text))`), lang)
		if err != nil {
			log.Fatalln("Could not construct query")
		}
		fmt.Printf("CaptureCount: %d\n PatternCount: %d\n StringCount: %d\n", q.CaptureCount(), q.PatternCount(), q.StringCount())

		// Execute query
		qc := sitter.NewQueryCursor()
		qc.Exec(q, n)

		// Check matches
		found := true
		for found {
			m, found := qc.NextMatch()
			if found {
				fmt.Println(m.ID, m.Captures)
			} else {
				break
			}
		}
		q.Close()
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
