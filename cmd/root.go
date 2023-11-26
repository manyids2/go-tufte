package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tufte",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", "./README.md", "Path to md file.")
	rootCmd.PersistentFlags().StringP("datadir", "d", "~/notes", "Path to notes.")
	rootCmd.PersistentFlags().StringP("htmldir", "s", "./assets", "Path to frontend.")
}
