// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package query

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Article struct {
	ID        pgtype.UUID
	Title     string
	Slug      string
	Body      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}