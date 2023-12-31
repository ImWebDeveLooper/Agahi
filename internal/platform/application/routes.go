package application

import v1 "agahi/internal/delivery/v1"

func (a *App) RegisterRoutes() {
	a.Router.POST("/register", v1.RegisterUserAction(a.Repository.UserRepo))
	a.Router.POST("/login", v1.LoginUserAction(a.Repository.UserRepo))
}
