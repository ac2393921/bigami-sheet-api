package repository

import (
	"context"

	"github.com/ac2393921/bigami-sheet-api/internal/domain/entity"
)

type ISchoolRepository interface {
	FetchAll(ctx context.Context) ([]*entity.School, error)
}
