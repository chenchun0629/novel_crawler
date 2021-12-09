package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Novel holds the schema definition for the Novel entity.
type Novel struct {
	ent.Schema
}

func (Novel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Novel.
func (Novel) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("小说名称").Annotations(entsql.Annotation{
			Size: 128,
		}),
		field.String("author").Comment("作者").Annotations(entsql.Annotation{
			Size: 64,
		}),
		field.String("category_title").Comment("小说类型"),
		field.String("intro").Comment("简介").Annotations(entsql.Annotation{
			Size: 1024,
		}),
	}
}

// Edges of the Novel.
func (Novel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clue", NovelClue.Type),
	}
}
