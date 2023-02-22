// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedDelete = `-- name: SchemaScopedDelete :exec
DELETE FROM foo.bar WHERE id = $1
`

func (q *Queries) SchemaScopedDelete(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, schemaScopedDelete, id)
	return err
}