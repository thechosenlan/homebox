// Code generated by ent, DO NOT EDIT.

package group

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeLocations holds the string denoting the locations edge name in mutations.
	EdgeLocations = "locations"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeLabels holds the string denoting the labels edge name in mutations.
	EdgeLabels = "labels"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// EdgeInvitationTokens holds the string denoting the invitation_tokens edge name in mutations.
	EdgeInvitationTokens = "invitation_tokens"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "group_users"
	// LocationsTable is the table that holds the locations relation/edge.
	LocationsTable = "locations"
	// LocationsInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationsInverseTable = "locations"
	// LocationsColumn is the table column denoting the locations relation/edge.
	LocationsColumn = "group_locations"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "group_items"
	// LabelsTable is the table that holds the labels relation/edge.
	LabelsTable = "labels"
	// LabelsInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelsInverseTable = "labels"
	// LabelsColumn is the table column denoting the labels relation/edge.
	LabelsColumn = "group_labels"
	// DocumentsTable is the table that holds the documents relation/edge.
	DocumentsTable = "documents"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// DocumentsColumn is the table column denoting the documents relation/edge.
	DocumentsColumn = "group_documents"
	// InvitationTokensTable is the table that holds the invitation_tokens relation/edge.
	InvitationTokensTable = "group_invitation_tokens"
	// InvitationTokensInverseTable is the table name for the GroupInvitationToken entity.
	// It exists in this package in order to avoid circular dependency with the "groupinvitationtoken" package.
	InvitationTokensInverseTable = "group_invitation_tokens"
	// InvitationTokensColumn is the table column denoting the invitation_tokens relation/edge.
	InvitationTokensColumn = "group_invitation_tokens"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCurrency,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Currency defines the type for the "currency" enum field.
type Currency string

// CurrencyChf is the default value of the Currency enum.
const DefaultCurrency = CurrencyChf

// Currency values.
const (
	CurrencyChf Currency = "chf"
	CurrencyUsd Currency = "usd"
	CurrencyEur Currency = "eur"
	CurrencyGbp Currency = "gbp"
	CurrencyJpy Currency = "jpy"
	CurrencyZar Currency = "zar"
	CurrencyAud Currency = "aud"
	CurrencyNok Currency = "nok"
	CurrencySek Currency = "sek"
	CurrencyDkk Currency = "dkk"
	CurrencyRmb Currency = "rmb"
	CurrencyBgn Currency = "bgn"
)

func (c Currency) String() string {
	return string(c)
}

// CurrencyValidator is a validator for the "currency" field enum values. It is called by the builders before save.
func CurrencyValidator(c Currency) error {
	switch c {
	case CurrencyChf, CurrencyUsd, CurrencyEur, CurrencyGbp, CurrencyJpy, CurrencyZar, CurrencyAud, CurrencyNok, CurrencySek, CurrencyDkk, CurrencyInr, CurrencyRmb, CurrencyBgn:
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for currency field: %q", c)
	}
}
