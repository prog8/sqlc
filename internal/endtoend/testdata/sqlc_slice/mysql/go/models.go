// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"database/sql"

	"github.com/prog8/sqlc-testdata/mysql"
)

type Foo struct {
	ID    int32
	Name  string
	Bar   sql.NullString
	Mystr mysql.ID
}
