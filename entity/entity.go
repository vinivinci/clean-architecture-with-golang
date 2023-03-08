package entity

import "errors"

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUserApplication() *User {
	return &User{}
}

func (u *User) IsValid() error {
	if u.ID < 1 {
		return errors.New("Users id must be >= 1")
	}
	if len(u.Name) > 200 {
		return errors.New("Length name must be < 200")
	}
	if len(u.Email) > 200 {
		return errors.New("Length email must be < 200")
	}
	return nil
}
