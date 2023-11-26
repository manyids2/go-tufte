package cmd

import (
	"log"
	"os"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
	"github.com/manyids2/go-tree-sitter-with-markdown/markdown"
	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cli",
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

		// Walk tree and print ( we are actually calling here )
		WalkWithIndent(n, 0, CallbackPrintNodeRange)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
