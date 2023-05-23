package application

import (
	"github.com/avinashb98/munshee/config"
	"github.com/avinashb98/munshee/datasources"
	"github.com/avinashb98/munshee/repository"
	"github.com/avinashb98/munshee/service"
	"log"
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
	c := config.Get()

	initialiseMongo(c.Mongo)

	//userRepo := repository.NewUserInmemory()
	userMongoRepo := repository.NewUserMongoRepository(c.Mongo)
	userSvc := service.NewUserService(userMongoRepo)

	//accountRepo := repository.NewAccountInmemory()
	accountMongoRepo := repository.NewAccountMongoRepository(c.Mongo)
	accountSvc := service.NewAccountService(accountMongoRepo, userSvc)

	tagRepo := repository.NewTagInmemory()
	tagSvc := service.NewTagService(tagRepo)

	txnRepo := repository.NewTxnInmemory()
	txnSvc := service.NewTxnService(txnRepo, userSvc, tagSvc, accountSvc)

	application := Application{
		Services: Services{
			User:    userSvc,
			Account: accountSvc,
			Txn:     txnSvc,
			Tag:     tagSvc,
		},
		Config: c,
	}
	return &application
}

func initialiseMongo(conf config.Mongo) {
	err := datasources.InitMongoORM(conf)
	if err != nil {
		log.Panic(err)
	}
}
