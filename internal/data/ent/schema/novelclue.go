package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NovelClue holds the schema definition for the NovelClue entity.
type NovelClue struct {
	ent.Schema
}

func (NovelClue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}

func (NovelClue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the NovelClue.
func (NovelClue) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("小说名称").Annotations(entsql.Annotation{
			Size: 128,
		}),
		field.Time("date").SchemaType(map[string]string{
			dialect.MySQL: "date",
		}).Comment("日期"),
		field.Int("score").Comment("热搜指数"),
		field.String("author").Comment("作者").Annotations(entsql.Annotation{
			Size: 64,
		}),
		field.String("category_title").Comment("小说类型"),
		field.String("intro").Comment("简介").Annotations(entsql.Annotation{
			Size: 1024,
		}),
		field.String("link").Comment("搜索连接").Annotations(entsql.Annotation{Size: 512}),
	}
}

// Edges of the NovelClue.
func (NovelClue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("novel", Novel.Type).Unique().Ref("clue"),
	}
}
