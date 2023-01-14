package repository

import (
	"context"
	"dangquang9a/go-location/internal/location"
	"dangquang9a/go-location/internal/models"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository location.LocationRepository
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	require.NoError(s.T(), err)

	s.repository = NewLocRepository(s.DB)
}

func (s *Suite) TestGetAllTodo() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "location" LIMIT 200`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "created_at", "created_by"}))

	_, err := s.repository.GetAllLocation(context.Background())

	require.NoError(s.T(), err)
}

func (s *Suite) TestGetTodosByUserId() {
	var id = uuid.New().String()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "location" WHERE "location"."created_by" = $1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "created_at", "created_by"}))

	_, err := s.repository.GetLocationsByUserID(context.Background(), id)
	require.NoError(s.T(), err)
}

func (s *Suite) TestCreateTodo() {

	user := &models.Location{
		Id:        uuid.New().String(),
		Name:      "test",
		CreatedAt: time.Now(),
		CreatedBy: uuid.New().String(),
	}

	s.mock.ExpectBegin()

	s.mock.ExpectExec(regexp.
		QuoteMeta(`INSERT INTO "location" ("id","content","created_at","created_by") VALUES ($1,$2,$3,$4)`)).
		WithArgs(user.Id, user.Name, user.CreatedAt, user.CreatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repository.CreateLocation(context.Background(), user)
	require.NoError(s.T(), err)
}
