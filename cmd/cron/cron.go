package cron

import (
	"context"
	"fmt"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/pkg/application"
	"github.com/novel_crawler/pkg/application/serve"
	"github.com/novel_crawler/pkg/conf"
	"github.com/spf13/cobra"
	"time"
)

var CronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务",
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

		var c = serve.NewXcron()
		_, _ = c.AddFunc("@every 1s", func() {
			fmt.Println("begin tick every 1 second")
			time.Sleep(5 * time.Second)
			fmt.Println("end tick every 1 second")
		})
		app.RegisterServes(
			c,
		)

		return app.Run()

	},
}
