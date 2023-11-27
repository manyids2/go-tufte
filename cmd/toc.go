package cmd

import (
	"log"

	"github.com/manyids2/go-tufte/core"
	"github.com/spf13/cobra"
)

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
		core.PrintToC(path)
	},
}

func init() {
	rootCmd.AddCommand(tocCmd)
}
