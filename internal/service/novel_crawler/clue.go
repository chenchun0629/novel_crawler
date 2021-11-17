package novel_crawler

import (
	"context"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/internal/data/ent"
	"time"
)

type Clue struct {
	Intro         string
	Link          string
	Score         int
	Title         string
	Date          time.Time
	Author        string
	CategoryTitle string
}

func (c Clue) Save(ctx context.Context) (*ent.NovelClue, error) {
	return data.DBClient.NovelClue.Create().
		SetIntro(c.Intro).
		SetLink(c.Link).
		SetScore(c.Score).
		SetTitle(c.Title).
		SetDate(c.Date).
		SetAuthor(c.Author).
		SetCategoryTitle(c.CategoryTitle).
		Save(ctx)
}
