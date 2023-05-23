package repository

import (
	"fmt"
	"github.com/avinashb98/munshee/config"
	"github.com/avinashb98/munshee/entity"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type userMongo struct {
	mongoConfig config.Mongo
}

func (u userMongo) CreateUser(username string, name string, email string) (*entity.User, error) {
	_, err := u.getByUsername(username)
	if err == nil {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}
	newUser := NewUserMongo(username, name, email)
	err = mgm.Coll(newUser).Create(newUser)
	return newUser.ToEntity(), err
}

func (u userMongo) getByUsername(username string) (*entity.User, error) {
	ctx := mgm.Ctx()
	foundUser := &UserMongo{}
	err := mgm.Coll(foundUser).FirstWithCtx(ctx, bson.M{"username": username}, foundUser)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return foundUser.ToEntity(), nil
}

func (u userMongo) Get(username string) (*entity.User, error) {
	return u.getByUsername(username)
}

func NewUserMongoRepository(mongoConfig config.Mongo) User {
	return &userMongo{
		mongoConfig: mongoConfig,
	}
}

type accountMongo struct {
	mongoConfig config.Mongo
}

func (a accountMongo) CreateAccount(username string, name string) (*entity.Account, error) {
	_, err := a.get(username, name)
	if err == nil {
		return nil, fmt.Errorf("account with username %s and name %s already exists", username, name)
	}

	newAccount := NewAccountMongo(username, name)
	err = mgm.Coll(newAccount).Create(newAccount)
	return newAccount.ToEntity(), err
}

func (a accountMongo) get(username string, name string) (*entity.Account, error) {
	ctx := mgm.Ctx()
	foundAccount := &AccountMongo{}
	err := mgm.Coll(foundAccount).FirstWithCtx(ctx, bson.M{"username": username, "name": name}, foundAccount)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}
	return foundAccount.ToEntity(), nil
}

func (a accountMongo) Get(username string, name string) (*entity.Account, error) {
	ctx := mgm.Ctx()
	foundAccount := &AccountMongo{}
	err := mgm.Coll(foundAccount).FirstWithCtx(ctx, bson.M{"username": username, "name": name}, foundAccount)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}
	return foundAccount.ToEntity(), nil
}

func (a accountMongo) GetAll(username string) ([]entity.Account, error) {
	ctx := mgm.Ctx()
	foundAccounts := make([]AccountMongo, 0)
	err := mgm.Coll(&AccountMongo{}).SimpleFindWithCtx(ctx, &foundAccounts, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	accounts := make([]entity.Account, 0)
	for _, account := range foundAccounts {
		accounts = append(accounts, *account.ToEntity())
	}
	return accounts, nil
}

func NewAccountMongoRepository(mongoConfig config.Mongo) Account {
	return &accountMongo{
		mongoConfig: mongoConfig,
	}
}
