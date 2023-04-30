package repository

import (
	"context"

	"github.com/ac2393921/bigami-sheet-api/internal/infrastructure/graphql/model"
)

type ISchoolRepository interface {
	FetchAll(ctx context.Context) ([]*model.School, error)
}
