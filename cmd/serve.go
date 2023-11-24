/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		htmldir, _ := cmd.Flags().GetString("htmldir")
		datadir, _ := cmd.Flags().GetString("datadir")
		fmt.Println("datadir: ", datadir)

		router := gin.Default()
		router.Static("/assets", htmldir)
		router.LoadHTMLFiles(htmldir + "/index.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"datadir": datadir})
		})

		router.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
