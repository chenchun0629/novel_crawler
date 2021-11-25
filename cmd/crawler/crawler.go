package crawler

import (
	"context"
	"github.com/novel_crawler/internal/biz/collector"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/pkg/application"
	"github.com/novel_crawler/pkg/application/serve"
	"github.com/novel_crawler/pkg/conf"
	"github.com/novel_crawler/pkg/log"
	"github.com/spf13/cobra"
)

var CrawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "爬取小说线索",

	Run: func(cmd *cobra.Command, args []string) {
		var app = application.NewApp(context.Background())
		app.RegisterBeforeAppStart(func(ctx context.Context) error {
			log.Init()
			conf.Init(cmd.Flag("config").Value.String(), "yaml")
			data.Init()
			return nil
		})

		app.RegisterAfterAppStart(func(ctx context.Context) error {
			log.Defer()
			return nil
		})

		app.RegisterServes(serve.NewCommand(func() error {
			collector.NewNovelClueCollector().Visit()

			//helper := log.NewHelper(log.DefaultLogger)
			//helper.Info("info message")
			//helper.Infof("info %s", "message")
			//helper.Infow("hello world", "key", "value")
			return nil
		}))

		var _ = app.Run()

		//
	},
	//PersistentPostRun: func(cmd *cobra.Command, args []string) {
	//	log.Defer()
	//},
}
