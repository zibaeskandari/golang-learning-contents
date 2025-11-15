package model

import (
	"context"
)

type OperatorRepository interface {
	Get(ctx context.Context) ([]Operator, error)
}
