// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"github.com/jackc/pgtype"
)

type TestTable struct {
	VBitNull    pgtype.Varbit
	VVarbitNull pgtype.Varbit
	VBit        pgtype.Varbit
	VVarbit     pgtype.Varbit
}
