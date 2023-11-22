// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package querytest

import (
	"context"
)

const demo = `-- name: Demo :one
SELECT CAST(GREATEST(1,2,3,4,5) AS UNSIGNED) as col1
`

func (q *Queries) Demo(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, demo)
	var col1 int64
	err := row.Scan(&col1)
	return col1, err
}
