package entity

import "fmt"

type AccountIn struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Account struct {
	ID       string
	Username string
	Name     string
	Balance  float64
}

func (a Account) ToOut() *AccountOut {
	return &AccountOut{
		Username: a.Username,
		Name:     a.Name,
		Balance:  fmt.Sprintf("%.2f", a.Balance),
	}
}

type AccountOut struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Balance  string `json:"balance"`
}
