package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Post     int
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	dateSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host,
		s.Post,
		s.UserName,
		s.Password,
		s.DbName,
	)

	s.Db = sqlx.MustConnect("postgres", dateSource)
	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error)
		return
	}
	fmt.Println("Connect Database OK")
}
func (s *Sql) Close() {

	s.Db.Close()
}
