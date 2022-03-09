package suite

import (
	"context"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"toki/code-golang/suite/model"
	mock_model "toki/code-golang/suite/model/.mocks"
	"xorm.io/core"
	"xorm.io/xorm"
)

type apiSuite struct {
	suite.Suite
	app App
	db  *xorm.Engine
}

func TestApi(t *testing.T) {
	a := &apiSuite{}
	suite.Run(t, a)
}

func (a *apiSuite) SetupSuite() {
	//dsn := os.Getenv(ENV_POSTGRES_DSN)
	//if dsn == "" {
	//	log.Fatalf("environment %s is empty", ENV_POSTGRES_DSN)
	//}
	db, err := xorm.NewEngine("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	a.Require().NoError(err)
	a.Require().NoError(db.Ping())
	db.SetMapper(core.GonicMapper{})
	db.ShowSQL(true)
	db.TZLocation = time.UTC
	db.DatabaseTZ = time.UTC

	// TODO: migrate
	_, err = db.Exec(`drop table if exists "user";`)
	a.Require().NoError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "user" (
		id bigserial not null primary key,
		name varchar(32) not null
	);`)
	a.Require().NoError(err)

	a.db = db
	a.app.storage = model.NewPostgresStorage(db)

}

func truncateDB(db *xorm.Session) error {
	beans := []interface{}{new(model.User)}
	for _, b := range beans {
		_, err := db.Where("1=1").Delete(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *apiSuite) TearDownSuite() {
	sess := s.db.NewSession()
	err := truncateDB(sess)
	s.Require().NoError(err)
	sess.Close()
	s.db.Close()
}

func (a *apiSuite) TestMockUser() {
	ctl := gomock.NewController(a.T())
	defer ctl.Finish()

	m := mock_model.NewMockStorage(ctl)
	m.EXPECT().GetUser(context.Background(), int64(1)).Return(model.User{
		ID:   1,
		Name: "Hello",
	}, nil).AnyTimes()

	user, err := m.GetUser(context.Background(), int64(1))
	a.Require().NoError(err)
	a.Require().Equal("Hello", user.Name)
}

func (a *apiSuite) TestUser() {
	id, err := a.app.storage.AddUser(model.User{Name: "Hello"})
	a.Require().NoError(err)

	user, err := a.app.storage.GetUser(context.Background(), id)
	a.Require().NoError(err)
	a.Require().Equal("Hello", user.Name)
}
