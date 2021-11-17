// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NovelCluesColumns holds the columns for the "novel_clues" table.
	NovelCluesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "score", Type: field.TypeInt},
		{Name: "author", Type: field.TypeString, Size: 64},
		{Name: "category_title", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString, Size: 1024},
		{Name: "link", Type: field.TypeString, Size: 512},
	}
	// NovelCluesTable holds the schema information for the "novel_clues" table.
	NovelCluesTable = &schema.Table{
		Name:       "novel_clues",
		Columns:    NovelCluesColumns,
		PrimaryKey: []*schema.Column{NovelCluesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NovelCluesTable,
	}
)

func init() {
	NovelCluesTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
}