package main

import (
	"github.com/fapodroper/internal/downloader"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "fapodroper"}

	var (
		username   string
		start, end int
	)

	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Specify the username")
	rootCmd.Flags().IntVar(&start, "start", 1, "Specify the start number")
	rootCmd.Flags().IntVar(&end, "end", 9999, "Specify the end number") // In fapodrop, the format `{username}_0001` is used in the name of the photos

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		downloader.DownloadImages(username, start, end)
	}

	rootCmd.Execute()
}
