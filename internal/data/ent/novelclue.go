// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/novel_crawler/internal/data/ent/novelclue"
)

// NovelClue is the model entity for the NovelClue schema.
type NovelClue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	// 小说名称
	Title string `json:"title,omitempty"`
	// Date holds the value of the "date" field.
	// 日期
	Date time.Time `json:"date,omitempty"`
	// Score holds the value of the "score" field.
	// 热搜指数
	Score int `json:"score,omitempty"`
	// Author holds the value of the "author" field.
	// 作者
	Author string `json:"author,omitempty"`
	// CategoryTitle holds the value of the "category_title" field.
	// 小说类型
	CategoryTitle string `json:"category_title,omitempty"`
	// Intro holds the value of the "intro" field.
	// 简介
	Intro string `json:"intro,omitempty"`
	// Link holds the value of the "link" field.
	// 搜索连接
	Link string `json:"link,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NovelClue) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case novelclue.FieldID, novelclue.FieldScore:
			values[i] = new(sql.NullInt64)
		case novelclue.FieldTitle, novelclue.FieldAuthor, novelclue.FieldCategoryTitle, novelclue.FieldIntro, novelclue.FieldLink:
			values[i] = new(sql.NullString)
		case novelclue.FieldCreateTime, novelclue.FieldUpdateTime, novelclue.FieldDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NovelClue", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NovelClue fields.
func (nc *NovelClue) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case novelclue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nc.ID = int(value.Int64)
		case novelclue.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				nc.CreateTime = value.Time
			}
		case novelclue.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				nc.UpdateTime = value.Time
			}
		case novelclue.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				nc.Title = value.String
			}
		case novelclue.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				nc.Date = value.Time
			}
		case novelclue.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				nc.Score = int(value.Int64)
			}
		case novelclue.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				nc.Author = value.String
			}
		case novelclue.FieldCategoryTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category_title", values[i])
			} else if value.Valid {
				nc.CategoryTitle = value.String
			}
		case novelclue.FieldIntro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intro", values[i])
			} else if value.Valid {
				nc.Intro = value.String
			}
		case novelclue.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				nc.Link = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this NovelClue.
// Note that you need to call NovelClue.Unwrap() before calling this method if this NovelClue
// was returned from a transaction, and the transaction was committed or rolled back.
func (nc *NovelClue) Update() *NovelClueUpdateOne {
	return (&NovelClueClient{config: nc.config}).UpdateOne(nc)
}

// Unwrap unwraps the NovelClue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nc *NovelClue) Unwrap() *NovelClue {
	tx, ok := nc.config.driver.(*txDriver)
	if !ok {
		panic("ent: NovelClue is not a transactional entity")
	}
	nc.config.driver = tx.drv
	return nc
}

// String implements the fmt.Stringer.
func (nc *NovelClue) String() string {
	var builder strings.Builder
	builder.WriteString("NovelClue(")
	builder.WriteString(fmt.Sprintf("id=%v", nc.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(nc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(nc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(nc.Title)
	builder.WriteString(", date=")
	builder.WriteString(nc.Date.Format(time.ANSIC))
	builder.WriteString(", score=")
	builder.WriteString(fmt.Sprintf("%v", nc.Score))
	builder.WriteString(", author=")
	builder.WriteString(nc.Author)
	builder.WriteString(", category_title=")
	builder.WriteString(nc.CategoryTitle)
	builder.WriteString(", intro=")
	builder.WriteString(nc.Intro)
	builder.WriteString(", link=")
	builder.WriteString(nc.Link)
	builder.WriteByte(')')
	return builder.String()
}

// NovelClues is a parsable slice of NovelClue.
type NovelClues []*NovelClue

func (nc NovelClues) config(cfg config) {
	for _i := range nc {
		nc[_i].config = cfg
	}
}
