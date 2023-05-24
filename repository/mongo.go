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

type txnMongo struct {
	mongoConfig config.Mongo
}

func (t txnMongo) CreateTxn(txnIn entity.TxnIn) (*entity.Txn, error) {
	newTxn := NewTxnMongo(txnIn)
	err := mgm.Coll(newTxn).Create(newTxn)
	return newTxn.ToEntity(), err
}

func (t txnMongo) Get(username string, id string) (*entity.Txn, error) {
	ctx := mgm.Ctx()
	foundTxn := &TxnMongo{}
	err := mgm.Coll(foundTxn).FirstWithCtx(ctx, bson.M{"username": username, "id": id}, foundTxn)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("txn not found")
		}
		return nil, err
	}
	return foundTxn.ToEntity(), nil
}

func (t txnMongo) GetAll(username string) ([]entity.Txn, error) {
	ctx := mgm.Ctx()
	foundTxns := make([]TxnMongo, 0)
	err := mgm.Coll(&TxnMongo{}).SimpleFindWithCtx(ctx, &foundTxns, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	txns := make([]entity.Txn, 0)
	for _, txn := range foundTxns {
		txns = append(txns, *txn.ToEntity())
	}
	return txns, nil
}

func (t txnMongo) UpdateTags(username string, id string, tags ...string) (*entity.Txn, error) {
	ctx := mgm.Ctx()
	foundTxn := &TxnMongo{}
	err := mgm.Coll(foundTxn).FirstWithCtx(ctx, bson.M{"username": username, "id": id}, foundTxn)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("txn not found")
		}
		return nil, err
	}
	foundTxn.Tags = tags
	err = mgm.Coll(foundTxn).Update(foundTxn)
	return foundTxn.ToEntity(), err
}

func NewTxnMongoRepository(mongoConfig config.Mongo) Txn {
	return &txnMongo{
		mongoConfig: mongoConfig,
	}
}

type tagMongo struct {
	mongoConfig config.Mongo
}

func (t tagMongo) CreateTag(name string) (*entity.Tag, error) {
	_, err := t.get(name)
	if err == nil {
		return nil, fmt.Errorf("tag with name %s already exists", name)
	}

	newTag := NewTagMongo(name)
	err = mgm.Coll(newTag).Create(newTag)
	return newTag.ToEntity(), err
}

func (t tagMongo) get(name string) (*entity.Tag, error) {
	ctx := mgm.Ctx()
	foundTag := &TagMongo{}
	err := mgm.Coll(foundTag).FirstWithCtx(ctx, bson.M{"name": name}, foundTag)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, fmt.Errorf("tag not found")
		}
		return nil, err
	}
	return foundTag.ToEntity(), nil
}

func (t tagMongo) UpsertTags(names ...string) ([]entity.Tag, error) {
	var uniqueTags []entity.Tag
	for _, name := range names {
		foundTag, err := t.get(name)
		if err == nil {
			uniqueTags = append(uniqueTags, *foundTag)
			continue
		}
		newTag := NewTagMongo(name)
		err = mgm.Coll(newTag).Create(newTag)
		if err != nil {
			return nil, err
		}
		uniqueTags = append(uniqueTags, *newTag.ToEntity())
	}
	return uniqueTags, nil
}

func (t tagMongo) GetAll() ([]entity.Tag, error) {
	ctx := mgm.Ctx()
	foundTags := make([]TagMongo, 0)
	err := mgm.Coll(&TagMongo{}).SimpleFindWithCtx(ctx, &foundTags, bson.M{})
	if err != nil {
		return nil, err
	}
	tags := make([]entity.Tag, 0)
	for _, tag := range foundTags {
		tags = append(tags, *tag.ToEntity())
	}
	return tags, nil
}

func NewTagMongoRepository(mongoConfig config.Mongo) Tag {
	return &tagMongo{
		mongoConfig: mongoConfig,
	}
}
