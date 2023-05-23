package application

import (
	"github.com/avinashb98/munshee/config"
	"github.com/avinashb98/munshee/repository"
	"github.com/avinashb98/munshee/service"
)

type Services struct {
	User    service.User
	Account service.Account
	Txn     service.Txn
	Tag     service.Tag
}

type Application struct {
	Services Services
	Config   config.Config
}

func Get() *Application {
	c := config.New()
	userRepo := repository.NewUserInmemory()
	userSvc := service.NewUserService(userRepo)

	accountRepo := repository.NewAccountInmemory()
	accountSvc := service.NewAccountService(accountRepo, userSvc)
	application := Application{
		Services: Services{
			User:    userSvc,
			Account: accountSvc,
		},
		Config: c,
	}
	return &application
}
