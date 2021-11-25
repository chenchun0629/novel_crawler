package crawler

import (
	"context"
	"github.com/novel_crawler/internal/biz/collector"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/pkg/application"
	"github.com/novel_crawler/pkg/application/serve"
	"github.com/novel_crawler/pkg/conf"
	"github.com/spf13/cobra"
)

var CrawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "爬取小说线索",

	RunE: func(cmd *cobra.Command, args []string) error {
		var app = application.NewApp(context.Background())
		app.RegisterBeforeAppStart(func(ctx context.Context) error {
			conf.Init(cmd.Flag("config").Value.String(), "yaml")
			data.Init()
			return nil
		})

		app.RegisterAfterAppStart(func(ctx context.Context) error {
			return nil
		})

		app.RegisterServes(
			serve.NewFunctional(collector.NewNovelClueCollector().Visit),
		)

		return app.Run()

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}
