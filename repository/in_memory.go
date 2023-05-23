package repository

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/google/uuid"
)

type usersStore struct {
	users map[string]entity.User
}

var us = usersStore{}

type userInMemory struct{}

func (u userInMemory) CreateUser(username string, name string, email string) (*entity.User, error) {
	_, found := us.users[username]
	if found {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}

	newUser := entity.User{
		Username: username,
		Name:     name,
		Email:    email,
		ID:       uuid.New().String(),
	}
	us.users[username] = newUser
	return &newUser, nil
}

func (u userInMemory) Get(username string) (*entity.User, error) {
	user, found := us.users[username]
	if !found {
		return nil, fmt.Errorf("user with username %s does not exist", username)
	}

	return &user, nil
}

func NewUserInmemory() User {
	return &userInMemory{}
}

type accountInmemory struct{}

func (a accountInmemory) GetBalance(id string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountInmemory) CreateAccount(id string, userID string, name string) (entity.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountInmemory) Get(id string) (entity.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountInmemory() Account {
	return &accountInmemory{}
}

type txnInmemory struct{}

func (t txnInmemory) CreateTxn(txn entity.Txn) (entity.Txn, error) {
	//TODO implement me
	panic("implement me")
}

func (t txnInmemory) Get(id string) (entity.Txn, error) {
	//TODO implement me
	panic("implement me")
}

func (t txnInmemory) GetAll(userID string) ([]entity.Txn, error) {
	//TODO implement me
	panic("implement me")
}

func (t txnInmemory) UpdateTags(id string, tags ...entity.Tag) (entity.Txn, error) {
	panic("implement me")
}

func NewTxnInmemory() Txn {
	return &txnInmemory{}
}

type tagInmemory struct{}

func (t tagInmemory) CreateTag(name string) (entity.Tag, error) {
	panic("implement me")
}

func NewTagInmemory() Tag {
	return &tagInmemory{}
}
