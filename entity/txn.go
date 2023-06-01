package entity

type TxnIn struct {
	Description string   `json:"description"`
	ToAccount   *string  `json:"to_account"`
	FromAccount *string  `json:"from_account"`
	Amount      float64  `json:"amount"`
	Tags        []string `json:"tags"`
	Username    string   `json:"username"`
	Emoji       string   `json:"emoji"`
}

type Txn struct {
	ID          string
	Description string
	ToAccount   *string
	FromAccount *string
	Username    string
	Amount      float64
	Tags        []string
	Emoji       string
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
		Emoji:       t.Emoji,
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
	Emoji       string   `json:"emoji"`
}
