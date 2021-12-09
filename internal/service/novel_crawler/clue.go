package novel_crawler

import (
	"context"
	"github.com/jinzhu/now"
	"github.com/novel_crawler/internal/data"
	"github.com/novel_crawler/internal/data/ent"
	"github.com/novel_crawler/internal/data/ent/novel"
	"github.com/novel_crawler/internal/data/ent/novelclue"
	"github.com/pkg/errors"
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

func (c Clue) Save(ctx context.Context) (*ent.Novel, error) {
	var (
		novelQuery = data.EntClient.Novel.Query()
		nd, err    = novelQuery.Where(
			novel.TitleEQ(c.Title),
		).First(ctx)
	)

	if err != nil && errors.Is(err, &ent.NotFoundError{}) {
		return nil, err
	}

	err = data.WithTx(ctx, data.EntClient, func(tx *ent.Tx) error {
		var err error
		if nd == nil {
			nd, err = data.EntClient.Novel.Create().
				SetIntro(c.Intro).
				SetTitle(c.Title).
				SetAuthor(c.Author).
				SetCategoryTitle(c.CategoryTitle).
				Save(ctx)
		} else {
			nd, err = nd.Update().
				SetIntro(c.Intro).
				SetTitle(c.Title).
				SetAuthor(c.Author).
				SetCategoryTitle(c.CategoryTitle).
				Save(ctx)
		}

		if err != nil {
			return err
		}

		var (
			query      = data.EntClient.NovelClue.Query()
			clue, err2 = query.Where(
				novelclue.AuthorEQ(c.Author),
				novelclue.TitleEQ(c.Title),
				novelclue.DateEQ(now.New(c.Date).BeginningOfDay()),
			).First(ctx)
		)

		if err2 != nil && errors.Is(err, &ent.NotFoundError{}) {
			return err2
		}

		if clue == nil {
			_, err = data.EntClient.NovelClue.Create().
				SetIntro(c.Intro).
				SetLink(c.Link).
				SetScore(c.Score).
				SetTitle(c.Title).
				SetDate(c.Date).
				SetAuthor(c.Author).
				SetCategoryTitle(c.CategoryTitle).
				SetNovel(nd).
				Save(ctx)
		} else {
			_, err = clue.Update().
				SetIntro(c.Intro).
				SetLink(c.Link).
				SetScore(c.Score).
				SetTitle(c.Title).
				SetDate(c.Date).
				SetAuthor(c.Author).
				SetCategoryTitle(c.CategoryTitle).
				SetNovel(nd).
				Save(ctx)
		}

		return err
	})

	return nd, err
}
