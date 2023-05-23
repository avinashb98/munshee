package entity

type CreateUserInput struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type User struct {
	ID        string
	Username  string
	Name      string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}
