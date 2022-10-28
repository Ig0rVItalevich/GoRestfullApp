package repository

import (
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user perfume.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, second_name, username, sex, password_hash) VALUES($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Firstname, user.Secondname, user.Username, user.Sex, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (perfume.User, error) {
	var user perfume.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 and password_hash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
