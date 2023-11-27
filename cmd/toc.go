package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manyids2/go-tufte/core"
	"github.com/spf13/cobra"
)

// Print table of contents to stdout
func PrintToC(path string) {
	doc, err := core.NewDocument(path)
	if err != nil {
		fmt.Println("go-tufte:", err)
		os.Exit(1)
	}

	doc.PrintSections(os.Stdout)
}

// tocCmd represents the toc command
var tocCmd = &cobra.Command{
	Use:   "toc",
	Short: "Print table of contents",
	Long: `Print table of contents with 
  - indent in tree ( skips -> one indent )
  - text
  - separator ( |> )
  - heading level
  - [startrow - endrow]
  - (startbyte - endbyte)

  e.g.
  go-tufte |> 1 [0, 3] (0, 58)
    Markdown specs |> 2 [4, 22] (59, 781)
  `,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Please specify --path.")
		}
		PrintToC(path)
	},
}

func init() {
	rootCmd.AddCommand(tocCmd)
}
