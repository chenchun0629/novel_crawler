package collector

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/novel_crawler/internal/data"
	"strconv"
	"strings"
	"time"
)

type NovelClueCollector struct {
	collector *colly.Collector
}

func NewNovelClueCollector() *NovelClueCollector {
	var c = colly.NewCollector()
	return &NovelClueCollector{collector: c}
}

func (c NovelClueCollector) Visit() {
	c.collector.OnHTML("div.category-wrap_iQLoo", func(e *colly.HTMLElement) {
		var (
			title  = e.ChildText("a[href]>div.c-single-text-ellipsis")
			link   = e.ChildAttr("a[href].title_dIF3B", "href")
			author = e.ChildTexts("div.intro_1l0wp")
			intro  = e.ChildText("div.c-single-text-ellipsis.desc_3CTjT")
			score  = e.ChildText("div.hot-index_1Bl1a")
		)

		fmt.Printf("%s, %s, %s, %s, %s \n\n\n", title, score, link, author, intro)

		var iScore, _ = strconv.Atoi(score)
		data.DBClient.NovelClue.Create().
			SetIntro(intro).
			SetLink(link).
			SetScore(iScore).
			SetTitle(title).
			SetDate(time.Now()).
			SetAuthor(strings.TrimPrefix(author[0], "作者：")).
			SetCategoryTitle(strings.TrimPrefix(author[1], "类型：")).
			SaveX(context.Background())
	})
	// todo error handler
	_ = c.collector.Visit("https://top.baidu.com/board?tab=novel")
}
