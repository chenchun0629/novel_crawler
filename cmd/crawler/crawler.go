package crawler

import (
	"github.com/novel_crawler/internal/biz/collector"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/pkg/conf"
	"github.com/spf13/cobra"
)

var CrawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "爬取小说线索",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		conf.Init(cmd.Flag("config").Value.String(), "yaml")
		data.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		collector.NewNovelClueCollector().Visit()
	},
}

func init() {
	//var (
	//	logger = log.NewStdLogger(os.Stdout)
	//	helper = log.NewHelper(logger)
	//)
	//helper = log.NewHelper(log.NewFilter(logger))

}
