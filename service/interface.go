package service

import "github.com/avinashb98/munshee/entity"

type User interface {
	CreateUser(username string, name string, email string) (*entity.UserOut, error)
	Get(username string) (*entity.UserOut, error)
}

type Account interface {
	CreateAccount(username string, name string) (*entity.AccountOut, error)
	Get(username string, name string) (*entity.AccountOut, error)
	GetAll(username string) ([]entity.AccountOut, error)
}

type Txn interface {
	CreateTxn(txnIn entity.TxnIn) (*entity.TxnOut, error)
	Get(username string, id string) (*entity.TxnOut, error)
	GetAll(username string) ([]entity.TxnOut, error)
	UpdateTags(username string, id string, tags ...string) (*entity.TxnOut, error)
}

type Tag interface {
	CreateTag(name string) (*entity.Tag, error)
	UpsertTags(names ...string) ([]entity.Tag, error)
	GetAll() ([]entity.Tag, error)
}
