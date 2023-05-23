package repository

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/google/uuid"
	"time"
)

type usersStore struct {
	users map[string]entity.User
}

var us = usersStore{
	users: make(map[string]entity.User),
}

type userInMemory struct{}

func (u userInMemory) CreateUser(username string, name string, email string) (*entity.User, error) {
	_, found := us.users[username]
	if found {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}

	now := time.Now().Unix()
	newUser := entity.User{
		Username:  username,
		Name:      name,
		Email:     email,
		ID:        uuid.New().String(),
		CreatedAt: now,
		UpdatedAt: now,
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

func (a accountInmemory) GetAll(username string) ([]entity.Account, error) {
	userAccounts, userFound := accountStore[username]

	if !userFound {
		return nil, fmt.Errorf("no account associated with user %s", username)
	}
	return userAccounts, nil
}

var accountStore = make(map[string][]entity.Account)

func (a accountInmemory) CreateAccount(username string, name string) (*entity.Account, error) {
	userAccounts, userFound := accountStore[username]

	if !userFound {
		accountStore[username] = make([]entity.Account, 0)
		userAccounts = accountStore[username]
	}

	for _, account := range userAccounts {
		if account.Name == name {
			return nil, fmt.Errorf("account with name %s already exists for user", name)
		}
	}

	newAccount := entity.Account{
		ID:       uuid.New().String(),
		Name:     name,
		Username: username,
		Balance:  0,
	}
	accountStore[username] = append(accountStore[username], newAccount)
	return &newAccount, nil
}

func (a accountInmemory) Get(username string, name string) (*entity.Account, error) {
	userAccounts, userFound := accountStore[username]

	if !userFound {
		return nil, fmt.Errorf("no account associated with user %s", username)
	}
	var account entity.Account
	accountFound := false
	for _, a := range userAccounts {
		if a.Name == name {
			accountFound = true
			account = a
			break
		}
	}
	if !accountFound {
		return nil, fmt.Errorf("no account with name %s associated with user %s", name, username)
	}
	return &account, nil
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
