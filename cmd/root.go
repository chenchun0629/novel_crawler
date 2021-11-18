package main

import (
	"fmt"
	"github.com/novel_crawler/cmd/crawler"
	"github.com/novel_crawler/pkg/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "novel_crawler",
}

var logCmd = &cobra.Command{
	Use: "log",
	Run: func(cmd *cobra.Command, args []string) {

		//// helper
		helper := log.NewHelper(log.DefaultLogger)
		helper.Log(log.LevelInfo, "key", "value")
		helper.Info("info message")
		helper.Infof("info %s", "message")
		helper.Infow("hello world", "key", "value")

	},
}

func init() {
	rootCmd.AddCommand(crawler.CrawlerCmd)
	rootCmd.AddCommand(logCmd)
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
