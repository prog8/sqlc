// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Memory struct {
	ID        pgtype.UUID
	VampireID pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Vampire struct {
	ID pgtype.UUID
}
