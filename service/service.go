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
	txnRepository  repository.Txn
	userService    User
	accountService Account
	tagService     Tag
}

func (t txn) CreateTxn(txnIn entity.TxnIn) (*entity.TxnOut, error) {
	_, err := t.userService.Get(txnIn.Username)
	if err != nil {
		return nil, err
	}
	_, err = t.accountService.Get(txnIn.Username, *txnIn.FromAccount)
	if err != nil {
		return nil, err
	}

	if txnIn.ToAccount != nil {
		_, err = t.accountService.Get(txnIn.Username, *txnIn.ToAccount)
		if err != nil {
			return nil, err
		}
	}

	txn, err := t.txnRepository.CreateTxn(txnIn)
	if err != nil {
		return nil, err
	}
	go func() {
		_, err := t.tagService.UpsertTags(txnIn.Tags...)
		if err != nil {
			fmt.Printf("Error while upserting tags %+v\n", err)
		}
	}()
	return txn.ToOut(), nil
}

func (t txn) Get(username string, id string) (*entity.TxnOut, error) {
	txn, err := t.txnRepository.Get(username, id)
	if err != nil {
		return nil, err
	}
	return txn.ToOut(), nil
}

func (t txn) GetAll(username string) ([]entity.TxnOut, error) {
	_, err := t.userService.Get(username)
	if err != nil {
		return nil, err
	}
	txn, err := t.txnRepository.GetAll(username)
	if err != nil {
		return nil, err
	}
	var txnOut []entity.TxnOut
	for _, txn := range txn {
		txnOut = append(txnOut, *txn.ToOut())
	}
	return txnOut, nil
}

func (t txn) UpdateTags(username string, id string, tags ...string) (*entity.TxnOut, error) {
	txn, err := t.txnRepository.UpdateTags(username, id, tags...)
	if err != nil {
		return nil, err
	}
	go func() {
		_, err := t.tagService.UpsertTags(tags...)
		if err != nil {
			fmt.Printf("Error while upserting tags %+v\n", err)
		}
	}()
	return txn.ToOut(), nil
}

func NewTxnService(txnRepository repository.Txn, userService User, tagService Tag, accountService Account) Txn {
	return &txn{txnRepository: txnRepository, userService: userService, tagService: tagService, accountService: accountService}
}

type tag struct {
	tagRepository repository.Tag
}

func (t tag) UpsertTags(names ...string) ([]entity.Tag, error) {
	return t.tagRepository.UpsertTags(names...)
}

func (t tag) GetAll() ([]entity.Tag, error) {
	return t.tagRepository.GetAll()
}

func (t tag) CreateTag(name string) (*entity.Tag, error) {
	return t.tagRepository.CreateTag(name)
}

func NewTagService(tagRepository repository.Tag) Tag {
	return &tag{tagRepository}
}
