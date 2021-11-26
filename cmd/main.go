package main

import (
	"github.com/novel_crawler/cmd/crawler"
	"github.com/novel_crawler/cmd/cron"
	"github.com/novel_crawler/pkg/log"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

func Execute() {
	log.Init()
	defer func() {
		log.Defer()
	}()

	if err := rootCmd.Execute(); err != nil {
		log.Errorw("run cmd error", "err", err)
	}
}

var rootCmd = &cobra.Command{
	Use: "novel_crawler",
}

func init() {
	rootCmd.AddCommand(crawler.CrawlerCmd)
	rootCmd.AddCommand(cron.CronCmd)
	rootCmd.PersistentFlags().StringP("config", "c", "./configs", "config path")
}
