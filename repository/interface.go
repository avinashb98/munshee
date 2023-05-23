package repository

import "github.com/avinashb98/munshee/entity"

type User interface {
	CreateUser(username string, name string, email string) (*entity.User, error)
	Get(username string) (*entity.User, error)
}

type Account interface {
	CreateAccount(username string, name string) (*entity.Account, error)
	Get(username string, name string) (*entity.Account, error)
	GetAll(username string) ([]entity.Account, error)
}

type Txn interface {
	CreateTxn(txn entity.TxnIn) (*entity.Txn, error)
	Get(username string, id string) (*entity.Txn, error)
	GetAll(username string) ([]entity.Txn, error)
	UpdateTags(username string, id string, tags ...string) (*entity.Txn, error)
}

type Tag interface {
	CreateTag(name string) (*entity.Tag, error)
	UpsertTags(names ...string) ([]entity.Tag, error)
	GetAll() ([]entity.Tag, error)
}
