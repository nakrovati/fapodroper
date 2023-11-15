package main

import (
	"github.com/fapodrop-downloader/internal/downloader"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "img-downloader"}

	var username string

	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Specify the username")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		downloader.DownloadImages(username)
	}

	rootCmd.Execute()
}
