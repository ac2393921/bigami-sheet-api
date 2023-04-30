package repository

import (
	"context"
	"log"

	db "github.com/ac2393921/bigami-sheet-api/internal/infrastructure/mysql/models"

	"strconv"

	"github.com/ac2393921/bigami-sheet-api/internal/infrastructure/graphql/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MySQLSchoolRepository struct {
	exec boil.ContextExecutor
}

func New(exec boil.ContextExecutor) *MySQLSchoolRepository {
	return &MySQLSchoolRepository{exec: exec}
}

func (m *MySQLSchoolRepository) FetchAll(ctx context.Context) ([]*model.School, error) {
	schools, err := db.Schools().All(ctx, m.exec)
	log.Printf("schools: %v", schools)
	if err != nil {
		return nil, err
	}
	return convertSchool(schools), nil
}

func convertSchool(schools db.SchoolSlice) []*model.School {
	var result []*model.School
	for _, s := range schools {
		result = append(result, &model.School{
			ID:    strconv.Itoa(s.ID),
			Name:  s.Name,
			Style: s.Style,
			Enemy: s.Enemy,
		})
	}
	return result
}
