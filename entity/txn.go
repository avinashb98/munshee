package entity

import "github.com/google/uuid"

type TxnIn struct {
	Description string   `json:"description"`
	ToAccount   *string  `json:"to_account"`
	FromAccount *string  `json:"from_account"`
	Amount      float64  `json:"amount"`
	Tags        []string `json:"tags"`
	Username    string   `json:"username"`
}

type Txn struct {
	ID          string
	Description string
	ToAccount   *string
	FromAccount *string
	Username    string
	Amount      float64
	Tags        []string
	CreatedAt   int64
	UpdatedAt   int64
}

func (t Txn) ToOut() *TxnOut {
	return &TxnOut{
		Description: t.Description,
		ToAccount:   t.ToAccount,
		FromAccount: t.FromAccount,
		Amount:      t.Amount,
		Tags:        t.Tags,
		Username:    t.Username,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

type TxnOut struct {
	Description string   `json:"description"`
	ToAccount   *string  `json:"to_account"`
	FromAccount *string  `json:"from_account"`
	Amount      float64  `json:"amount"`
	Tags        []string `json:"tags"`
	Username    string   `json:"username"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
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
		var tagsStr []string
		for _, tag := range tags {
			tagsStr = append(tagsStr, tag.Name)
		}
		t.Tags = tagsStr
	}
}

func WithAmount(amount float64) TxnOption {
	return func(t *Txn) {
		t.Amount = amount
	}
}

func NewTxn(userID string, opts ...TxnOption) Txn {
	txn := Txn{Username: userID, ID: uuid.New().String()}

	for _, opt := range opts {
		opt(&txn)
	}
	return txn
}
