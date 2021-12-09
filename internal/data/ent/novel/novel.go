// Code generated by entc, DO NOT EDIT.

package novel

import (
	"time"
)

const (
	// Label holds the string label denoting the novel type in the database.
	Label = "novel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldCategoryTitle holds the string denoting the category_title field in the database.
	FieldCategoryTitle = "category_title"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// EdgeClue holds the string denoting the clue edge name in mutations.
	EdgeClue = "clue"
	// Table holds the table name of the novel in the database.
	Table = "novels"
	// ClueTable is the table that holds the clue relation/edge.
	ClueTable = "novel_clues"
	// ClueInverseTable is the table name for the NovelClue entity.
	// It exists in this package in order to avoid circular dependency with the "novelclue" package.
	ClueInverseTable = "novel_clues"
	// ClueColumn is the table column denoting the clue relation/edge.
	ClueColumn = "novel_clue"
)

// Columns holds all SQL columns for novel fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldAuthor,
	FieldCategoryTitle,
	FieldIntro,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)