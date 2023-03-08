package fixture

import (
	"database/sql"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: db}
}

func (ur *UserRepositoryDB) Insert(id int, name string, email string) error {
	smt, err := ur.db.Prepare(`
		Insert into users (id, name, email) values(?,?,?)
	`)
	if err != nil {
		return err
	}

	_, err = smt.Exec(
		id,
		name,
		email,
	)
	if err != nil {
		return err
	}
	return nil
}
