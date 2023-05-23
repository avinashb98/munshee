package entity

type UserIn struct {
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

func (u *User) ToOut() *UserOut {
	return &UserOut{
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserOut struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
