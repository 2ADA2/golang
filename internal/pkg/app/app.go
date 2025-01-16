package app

import (
	"fmt"
	"log"
	"myapp/internal/app/controllers"
	getPosts "myapp/internal/app/endpoint"
	"myapp/internal/app/getUsers"
	"myapp/internal/app/login"
	"myapp/internal/app/ping"
	"myapp/internal/app/service"
	"myapp/internal/mw/checkToken"
	"os"

	"github.com/labstack/echo/v4"
)

type App struct {
	e          *getPosts.GetPosts
	s          *service.Service
	ping       *ping.Ping
	getUsers   *getUsers.GetUsers
	createUser *controllers.CreateUser
	login      *login.Login

	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = getPosts.New(a.s)
	a.ping = ping.New()
	a.getUsers = getUsers.New()
	a.createUser = controllers.New()
	a.login = login.New()

	a.echo = echo.New()

	a.echo.GET("/ping", a.ping.Status)
	a.echo.POST("/register", a.createUser.Status)
	a.echo.POST("/login", a.login.Status)

	a.echo.GET("/getPosts", a.e.Status, checkToken.CheckToken)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
