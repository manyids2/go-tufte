package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manyids2/go-tufte/components"
	"github.com/manyids2/go-tufte/markdown"
	"github.com/spf13/cobra"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Interactive markdown viewer.",
	Long:  `Interactive markdown viewer`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Please specify --path.")
		}

		// Load document
		doc, err := markdown.NewDocument(path)
		if err != nil {
			fmt.Println("go-tufte:", err)
			os.Exit(1)
		}

		// Start app
		app := components.NewApp(doc)

		// Run the application
		err = app.Application.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
