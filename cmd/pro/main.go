package main

import (
	"fmt"
	"github-trending/db"
	"github-trending/handler"
	"github-trending/helper"
	log "github-trending/log"
	"github-trending/repository/repo_impl"
	"github-trending/router"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func init() {
	fmt.Println("init package name")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)

}
func main() {

	sql := &db.Sql{
		Host:     "35.222.222.36",
		Post:     5432,
		UserName: "thongvo",
		Password: "thongvo109",
		DbName:   "thongvo",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()
	go scheduleUpdateTrending(600* time.Second,repoHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
func scheduleUpdateTrending(timeSchedule time.Duration,handler handler.RepoHandler)  {
	ticker := time.NewTicker(timeSchedule)

	go func() {
		for  {
			select {
			case <- ticker.C:
				fmt.Println("Checking from Github")
				helper.CrawlRepo(handler.GithubRepo)

			}
		}
	}()

}
