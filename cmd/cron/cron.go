package cron

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

var CronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		conf.Init(cmd.Flag("config").Value.String(), "yaml")
		data.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var app = application.NewApp(context.Background())
		app.RegisterBeforeAppStart(func(ctx context.Context) error {
			return nil
		})

		app.RegisterAfterAppStart(func(ctx context.Context) error {
			return nil
		})

		type Config struct {
			Source string `json:"source" yaml:"source"`
		}

		var (
			c    = serve.NewXcron()
			_, _ = c.AddFunc("45 10 * * *", func() {
				if err := collector.NewNovelClueCollector().Visit(); err != nil {
					log.Errorw("执行任务采集小说线索任务失败", "err", err)
				} else {
					log.Info("执行任务采集小说线索任务完成")
				}
			})
		)

		//fmt.Println("================err", err)

		//_, _ = c.AddFunc("@every 1s", func() {
		//	fmt.Println("begin tick every 1 second")
		//	time.Sleep(5 * time.Second)
		//	fmt.Println("end tick every 1 second")
		//})
		app.RegisterServes(
			c,
		)

		return app.Run()

	},
}
