package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/manyids2/go-tufte/components"
	"github.com/manyids2/go-tufte/core"
	"github.com/spf13/cobra"
)

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Markdown viewer.",
	Long:  `Markdown viewer`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln("Please specify --path.")
		}

		// Load document
		doc, err := core.NewDocument(path)
		if err != nil {
			fmt.Println("go-tufte:", err)
			os.Exit(1)
		}
		// doc.PrintSections()

		// Start app
		app := components.NewApp()

		// Try to print rendered string
		content := string(*doc.Buffer)

		// Try using glamour to rener
		out, err := glamour.Render(content, "dark")

		// Parse ansi using tview
		// NOTE: Not working - why???
		// w := tview.ANSIWriter(app.Content)
		w := app.Content
		fmt.Fprintf(w, out)

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
