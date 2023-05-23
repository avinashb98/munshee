package service

import (
	"github.com/avinashb98/munshee/entity"
	"github.com/avinashb98/munshee/repository"
)

type account struct {
	accountRepository repository.Account
}

func (a account) GetBalance(id string) (float64, error) {
	return a.accountRepository.GetBalance(id)
}

func (a account) CreateAccount(id string, userID string, name string) (entity.Account, error) {
	return a.accountRepository.CreateAccount(id, userID, name)
}

func (a account) Get(id string) (entity.Account, error) {
	return a.accountRepository.Get(id)
}

func NewAccountService(accountRepository repository.Account) Account {
	return &account{accountRepository}
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

func (u user) Get(id string) (*entity.UserOut, error) {
	user, err := u.userRepository.Get(id)
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
