package entities

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *User) IsSuper() bool {
	return u.Role == "super"
}

func (u *User) IsMedium() bool {
	return u.Role == "medium"
}
func (u *User) IsRegular() bool {
	return u.Role == "regular"
}
