package usecase

import (
	"context"

	"github.com/ac2393921/bigami-sheet-api/internal/domain/entity"
	"github.com/ac2393921/bigami-sheet-api/internal/domain/repository"
)

type FetchAllSchoolsUsecase struct {
	schoolRepo repository.ISchoolRepository
}

func NewFetchAllSchoolsUsecase(schoolRepo repository.ISchoolRepository) *FetchAllSchoolsUsecase {
	return &FetchAllSchoolsUsecase{
		schoolRepo: schoolRepo,
	}
}

func (f *FetchAllSchoolsUsecase) Handle(ctx context.Context) ([]*entity.School, error) {
	schools, err := f.schoolRepo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}