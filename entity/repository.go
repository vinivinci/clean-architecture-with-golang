package entity

type UserRepository interface {
	Insert(id int, name string, email string) error
}
