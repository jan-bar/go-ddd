package cmd

import service "github.com/jettjia/gin-ddd/application/service/user"

type App struct {
	UserSrv *service.UserService
}

func NewApp(userSrv *service.UserService) *App {
	return &App{UserSrv: userSrv}
}
