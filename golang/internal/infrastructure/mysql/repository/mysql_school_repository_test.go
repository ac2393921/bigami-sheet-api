package repository

import (
	"context"
	"database/sql/driver"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ac2393921/bigami-sheet-api/internal/infrastructure/graphql/model"
	db "github.com/ac2393921/bigami-sheet-api/internal/infrastructure/mysql/models"
	"github.com/stretchr/testify/assert"
)

// type mockExecutor struct{}

// func (m *mockExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	return nil, nil
// }

func TestMySQLSchoolRepository_FetchAll(t *testing.T) {
	columns := []string{"id", "name", "style", "enemy"}
	query := "SELECT id, name, style, enemy FROM schools"
	
	ctx := context.Background()
	// mockExec := &mockExecutor{}
	// conn, err := sql.Open("mysql", "bigami:bigami@tcp(db)/bigami")
	exec, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	mockRow := []driver.Value{1, "School A",  "Style A", "Enemy A"}
	rows := sqlmock.NewRows(columns).AddRow(mockRow...)
	mock.ExpectQuery(regexp.QuoteMeta(query)).
				WillReturnRows(rows)

	mysqlRepo := New(exec)

	schools := db.SchoolSlice{
		{ID: 1, Name: "School A", Style: "Style A", Enemy: "Enemy A"},
		// {ID: 2, Name: "School B", Style: "Style B", Enemy: "Enemy B"},
	}

	// モックの変換関数を作成
	convertSchoolsMock := func (schools db.SchoolSlice) []*model.School {
		result := make([]*model.School, len(schools))
		for i, school := range schools {
			result[i] = &model.School{
				ID:      strconv.Itoa(school.ID),
				Name:    school.Name,
				Style:  school.Style,
				Enemy: school.Enemy,
			}
		}
		return result
	}

	excpeted := convertSchoolsMock(schools)

	result, err := mysqlRepo.FetchAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, excpeted, result)
}