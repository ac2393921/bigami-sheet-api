package contoller

import (
	"context"

	"github.com/ac2393921/bigami-sheet-api/internal/domain/entity"
	"github.com/ac2393921/bigami-sheet-api/internal/usecase"
)

type SchoolController struct {
	FetchAllSchoolsUsecase usecase.FetchAllSchoolsUsecase
}

func NewSchoolController(fetchAllSchoolsUsecase usecase.FetchAllSchoolsUsecase) *SchoolController {
	return &SchoolController{
		FetchAllSchoolsUsecase: fetchAllSchoolsUsecase,
	}
}

func (s *SchoolController) FetchAllSchools(ctx context.Context) ([]*entity.School, error) {
	response, err := s.FetchAllSchoolsUsecase.Handle(ctx)
	if err != nil {
		return nil, err
	}
	return response, err
}