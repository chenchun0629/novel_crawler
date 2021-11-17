package main

import (
	"fmt"
	"github.com/novel_crawler/cmd/crawler"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "novel_crawler",
}

func init() {
	rootCmd.AddCommand(crawler.CrawlerCmd)
	rootCmd.PersistentFlags().StringP("config", "c", "./configs", "config path")
}

func main() {
	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
