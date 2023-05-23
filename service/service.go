package service

import (
	"fmt"
	"github.com/avinashb98/munshee/entity"
	"github.com/avinashb98/munshee/repository"
)

type account struct {
	accountRepository repository.Account
	userService       User
}

func (a account) CreateAccount(username string, name string) (*entity.AccountOut, error) {
	_, err := a.userService.Get(username)
	if err != nil {
		return nil, err
	}
	account, err := a.accountRepository.CreateAccount(username, name)
	if err != nil {
		return nil, err
	}
	return account.ToOut(), nil
}

func (a account) Get(username string, name string) (*entity.AccountOut, error) {
	_, err := a.userService.Get(username)
	if err != nil {
		return nil, err
	}
	account, err := a.accountRepository.Get(username, name)
	fmt.Printf("%+v\n", account)
	if err != nil {
		return nil, err
	}
	return account.ToOut(), nil
}

func (a account) GetAll(username string) ([]entity.AccountOut, error) {
	_, err := a.userService.Get(username)
	if err != nil {
		return nil, err
	}
	accounts, err := a.accountRepository.GetAll(username)
	if err != nil {
		return nil, err
	}
	var accountsOut []entity.AccountOut
	for _, account := range accounts {
		accountsOut = append(accountsOut, *account.ToOut())
	}
	return accountsOut, nil
}

func NewAccountService(accountRepository repository.Account, userSvc User) Account {
	return &account{accountRepository: accountRepository, userService: userSvc}
}

type user struct {
	userRepository repository.User
}

func (u user) CreateUser(username string, name string, email string) (*entity.UserOut, error) {

	user, err := u.userRepository.CreateUser(username, name, email)
	if err != nil {
		return nil, err
	}
	return user.ToOut(), nil
}

func (u user) Get(username string) (*entity.UserOut, error) {
	user, err := u.userRepository.Get(username)
	if err != nil {
		return nil, err
	}
	return user.ToOut(), nil
}

func NewUserService(userRepository repository.User) User {
	return &user{userRepository}
}

type txn struct {
	txnRepository repository.Txn
}

func (t txn) CreateTxn(txn entity.Txn) (entity.Txn, error) {
	return t.txnRepository.CreateTxn(txn)
}

func (t txn) Get(id string) (entity.Txn, error) {
	return t.txnRepository.Get(id)
}

func (t txn) GetAll(userID string) ([]entity.Txn, error) {
	return t.txnRepository.GetAll(userID)
}

func (t txn) UpdateTags(id string, tags ...entity.Tag) (entity.Txn, error) {
	return t.txnRepository.UpdateTags(id, tags...)
}

func NewTxnService(txnRepository repository.Txn) Txn {
	return &txn{txnRepository}
}

type tag struct {
	tagRepository repository.Tag
}

func (t tag) CreateTag(name string) (entity.Tag, error) {
	return t.tagRepository.CreateTag(name)
}

func NewTagService(tagRepository repository.Tag) Tag {
	return &tag{tagRepository}
}
