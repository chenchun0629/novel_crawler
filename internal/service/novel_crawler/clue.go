package novel_crawler

import (
	"context"
	"errors"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/internal/data/ent"
	"github.com/novel_crawler/internal/data/ent/novelclue"
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
	var (
		query     = data.EntClient.NovelClue.Query()
		clue, err = query.Where(
			novelclue.AuthorEQ(c.Author),
			novelclue.TitleEQ(c.Title),
			novelclue.DateEQ(c.Date),
		).First(ctx)
	)

	var nfe *ent.NotFoundError
	if err != nil && !errors.As(err, &nfe) {
		return clue, err
	}

	if clue == nil {
		return data.EntClient.NovelClue.Create().
			SetIntro(c.Intro).
			SetLink(c.Link).
			SetScore(c.Score).
			SetTitle(c.Title).
			SetDate(c.Date).
			SetAuthor(c.Author).
			SetCategoryTitle(c.CategoryTitle).
			Save(ctx)
	} else {
		return clue.Update().
			SetIntro(c.Intro).
			SetLink(c.Link).
			SetScore(c.Score).
			SetTitle(c.Title).
			SetDate(c.Date).
			SetAuthor(c.Author).
			SetCategoryTitle(c.CategoryTitle).
			Save(ctx)
	}
}
