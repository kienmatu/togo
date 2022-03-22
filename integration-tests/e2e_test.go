package integrations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kienmatu/go-todos/config"
	"kienmatu/go-todos/internal/auth/presenter"
	"kienmatu/go-todos/internal/server"
	"kienmatu/go-todos/utils"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	config *config.Configuration
	dbConn *gorm.DB
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.config = &config.Configuration{
		Port:                  "8081",
		SigningKey:            "AA",
		HashSalt:              "AAA",
		DatabaseConnectionURL: "host=localhost user=postgres password=password1 dbname=todos port=5432",
		TokenTTL:              86400,
		JwtSecret:             "abcaa",
	}
	dsn := s.config.DatabaseConnectionURL
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	s.Require().NoError(err)
	s.dbConn = dbConn

	// s.dbMigration, err = migrate.New("file://../db/migration", s.dbConnectionStr)
	// s.Require().NoError(err)
	// if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
	// 	s.Require().NoError(err)
	// }
	serverReady := make(chan bool)
	server := server.NewServer(s.config, s.dbConn, logrus.New(), serverReady)

	go server.Run()
	<-serverReady
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

// func (s *e2eTestSuite) SetupTest() {
// 	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
// 		s.Require().NoError(err)
// 	}
// }

// func (s *e2eTestSuite) TearDownTest() {
// 	s.NoError(s.dbMigration.Down())
// }

func (s *e2eTestSuite) Test_EndToEnd_Register() {
	username := utils.RandString(10)
	pwd := utils.RandString(8)

	reqStr := `{"username":"` + username + `", "password": "` + pwd + `", "limit": 10}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%s/api/v1/auth/register", s.config.Port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	expectedResp := `{"username":"` + strings.ToLower(username) + `","limit":10}`
	var expectedUser = presenter.SignUpResponse{}
	err = json.Unmarshal([]byte(expectedResp), &expectedUser)
	s.NoError(err)

	var actualUser = presenter.SignUpResponse{}
	err = json.Unmarshal([]byte(strings.Trim(string(byteBody), "\n")), &actualUser)
	s.NoError(err)

	s.Equal(actualUser.Username, expectedUser.Username)
	s.Equal(actualUser.Limit, expectedUser.Limit)
	response.Body.Close()
}
