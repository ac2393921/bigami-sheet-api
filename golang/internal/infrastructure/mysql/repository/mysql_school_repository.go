package repository

import (
	"context"
	"log"

	"github.com/ac2393921/bigami-sheet-api/internal/domain/entity"
	db "github.com/ac2393921/bigami-sheet-api/internal/infrastructure/mysql/models"

	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MySQLSchoolRepository struct {
	exec boil.ContextExecutor
}

func NewSchoolRepository(exec boil.ContextExecutor) *MySQLSchoolRepository {
	return &MySQLSchoolRepository{exec: exec}
}

func (m *MySQLSchoolRepository) FetchAll(ctx context.Context) ([]*entity.School, error) {
	schools, err := db.Schools().All(ctx, m.exec)
	log.Printf("schools: %v", schools)
	if err != nil {
		return nil, err
	}
	return convertSchools(schools), nil
}

// SQLboilerをEntityに変換する
func convertSchools(schools db.SchoolSlice) []*entity.School {
	var result []*entity.School
	for _, s := range schools {
		result = append(result, &entity.School{
			ID:    strconv.Itoa(s.ID),
			Name:  s.Name,
			Style: s.Style,
			Enemy: s.Enemy,
		})
	}
	return result
}
