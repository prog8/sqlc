// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package querytest

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateMemories(ctx context.Context, vampireID []uuid.UUID) ([]Memory, error)
	GetVampireIDs(ctx context.Context, vampireID []uuid.UUID) ([]uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
