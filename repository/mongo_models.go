package repository

import (
	"github.com/avinashb98/munshee/entity"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"time"
)

type UserMongo struct {
	mgm.DefaultModel `bson:",inline"`

	ID        string `bson:"id"`
	Username  string `bson:"username"`
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (u *UserMongo) ToEntity() *entity.User {
	return &entity.User{
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUserMongo(username string, name string, email string) *UserMongo {
	return &UserMongo{
		ID:        uuid.New().String(),
		Username:  username,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

type AccountMongo struct {
	mgm.DefaultModel `bson:",inline"`

	ID       string  `bson:"id"`
	Username string  `bson:"username"`
	Name     string  `bson:"name"`
	Balance  float64 `bson:"balance"`
}

func (a *AccountMongo) ToEntity() *entity.Account {
	return &entity.Account{
		ID:       a.ID,
		Username: a.Username,
		Name:     a.Name,
		Balance:  a.Balance,
	}
}

func NewAccountMongo(username string, name string) *AccountMongo {
	return &AccountMongo{
		ID:       uuid.New().String(),
		Username: username,
		Name:     name,
		Balance:  0,
	}
}

type TagMongo struct {
	mgm.DefaultModel `bson:",inline"`

	ID   string `bson:"id"`
	Name string `bson:"name"`
}

func (t *TagMongo) ToEntity() *entity.Tag {
	return &entity.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func NewTagMongo(name string) *TagMongo {
	return &TagMongo{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type TxnMongo struct {
	mgm.DefaultModel `bson:",inline"`

	ID          string   `bson:"id"`
	Description string   `bson:"description"`
	Username    string   `bson:"username"`
	ToAccount   *string  `bson:"to_account"`
	FromAccount *string  `bson:"from_account"`
	Amount      float64  `bson:"amount"`
	Tags        []string `bson:"tags"`
	CreatedAt   int64    `bson:"created_at"`
	UpdatedAt   int64    `bson:"updated_at"`
}

func (t *TxnMongo) ToEntity() *entity.Txn {
	return &entity.Txn{
		ID:          t.ID,
		Description: t.Description,
		Username:    t.Username,
		ToAccount:   t.ToAccount,
		FromAccount: t.FromAccount,
		Amount:      t.Amount,
		Tags:        t.Tags,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func NewTxnMongo(txnIn entity.TxnIn) *TxnMongo {
	return &TxnMongo{
		ID:          uuid.New().String(),
		Description: txnIn.Description,
		Username:    txnIn.Username,
		ToAccount:   txnIn.ToAccount,
		FromAccount: txnIn.FromAccount,
		Amount:      txnIn.Amount,
		Tags:        txnIn.Tags,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}
