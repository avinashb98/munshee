package entity

import "github.com/google/uuid"

type Txn struct {
	ID          string
	Description string
	ToAccount   *string
	FromAccount *string
	UserID      string
	Amount      float64
	Tag         []Tag
	Timestamp   int64
}

type TxnOption func(*Txn)

func WithDescription(description string) TxnOption {
	return func(t *Txn) {
		t.Description = description
	}
}

func WithToAccount(toAccount *string) TxnOption {
	return func(t *Txn) {
		t.ToAccount = toAccount
	}
}

func WithFromAccount(fromAccount *string) TxnOption {
	return func(t *Txn) {
		t.FromAccount = fromAccount
	}
}

func WithTags(tags ...Tag) TxnOption {
	return func(t *Txn) {
		t.Tag = tags
	}
}

func WithAmount(amount float64) TxnOption {
	return func(t *Txn) {
		t.Amount = amount
	}
}

func NewTxn(userID string, opts ...TxnOption) Txn {
	txn := Txn{UserID: userID, ID: uuid.New().String()}

	for _, opt := range opts {
		opt(&txn)
	}
	return txn
}
