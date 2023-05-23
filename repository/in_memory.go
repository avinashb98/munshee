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

var txnStore = make(map[string][]entity.Txn)

func (t txnInmemory) CreateTxn(txnIn entity.TxnIn) (*entity.Txn, error) {
	_, userFound := accountStore[txnIn.Username]
	if !userFound {
		txnStore[txnIn.Username] = make([]entity.Txn, 0)
	}

	newTxn := entity.Txn{
		ID:          uuid.New().String(),
		Username:    txnIn.Username,
		FromAccount: txnIn.FromAccount,
		ToAccount:   txnIn.ToAccount,
		Amount:      txnIn.Amount,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		Description: txnIn.Description,
		Tags:        txnIn.Tags,
	}
	txnStore[txnIn.Username] = append(txnStore[txnIn.Username], newTxn)
	return &newTxn, nil
}

func (t txnInmemory) Get(username string, id string) (*entity.Txn, error) {
	txns, found := txnStore[username]
	if !found {
		return nil, fmt.Errorf("no txn associated with user %s", username)
	}
	var txn entity.Txn
	txnFound := false
	for _, t := range txns {
		if t.ID == id {
			txnFound = true
			txn = t
			break
		}
	}
	if !txnFound {
		return nil, fmt.Errorf("no txn with id %s associated with user %s", id, username)
	}
	return &txn, nil
}

func (t txnInmemory) GetAll(username string) ([]entity.Txn, error) {
	txns, found := txnStore[username]
	if !found {
		return nil, fmt.Errorf("no txn associated with user %s", username)
	}
	return txns, nil
}

func (t txnInmemory) UpdateTags(username string, id string, tags ...string) (*entity.Txn, error) {
	txns, found := txnStore[username]
	if !found {
		return nil, fmt.Errorf("no txn associated with user %s", username)
	}
	var txn entity.Txn
	txnFound := false
	for _, t := range txns {
		if t.ID == id {
			txnFound = true
			txn = t
			break
		}
	}
	if !txnFound {
		return nil, fmt.Errorf("no txn with id %s associated with user %s", id, username)
	}
	var tagsStr []string
	for _, tag := range tags {
		tagsStr = append(tagsStr, tag)
	}
	txn.Tags = tagsStr
	return &txn, nil
}

func NewTxnInmemory() Txn {
	return &txnInmemory{}
}

type tagInmemory struct{}

var tagStore = make([]entity.Tag, 0)

func (t tagInmemory) CreateTag(name string) (*entity.Tag, error) {
	for _, tag := range tagStore {
		if tag.Name == name {
			return nil, fmt.Errorf("tag with name %s already exists", name)
		}
	}
	newTag := entity.Tag{
		ID:   uuid.New().String(),
		Name: name,
	}
	tagStore = append(tagStore, newTag)
	return &newTag, nil
}

func (t tagInmemory) get(name string) (*entity.Tag, bool) {

	for _, tag := range tagStore {
		if tag.Name == name {
			return &tag, true
		}
	}
	return nil, false
}

func (t tagInmemory) UpsertTags(names ...string) ([]entity.Tag, error) {
	var uniqueTags []entity.Tag
	for _, name := range names {
		tag, found := t.get(name)
		var err error
		if !found {
			tag, err = t.CreateTag(name)
			if err != nil {
				return nil, err
			}
		}
		uniqueTags = append(uniqueTags, *tag)
	}
	return uniqueTags, nil
}

func (t tagInmemory) GetAll() ([]entity.Tag, error) {
	return tagStore, nil
}

func NewTagInmemory() Tag {
	return &tagInmemory{}
}
