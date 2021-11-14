// Code generated by entc, DO NOT EDIT.

package userfrozen

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userfrozen type in the database.
	Label = "user_frozen"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldFrozenBy holds the string denoting the frozen_by field in the database.
	FieldFrozenBy = "frozen_by"
	// FieldFrozenCause holds the string denoting the frozen_cause field in the database.
	FieldFrozenCause = "frozen_cause"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUnfrozenBy holds the string denoting the unfrozen_by field in the database.
	FieldUnfrozenBy = "unfrozen_by"
	// Table holds the table name of the userfrozen in the database.
	Table = "user_frozens"
)

// Columns holds all SQL columns for userfrozen fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldFrozenBy,
	FieldFrozenCause,
	FieldCreateAt,
	FieldEndAt,
	FieldStatus,
	FieldUnfrozenBy,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
