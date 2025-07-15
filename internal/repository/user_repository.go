package repository

import (
	"dballz/internal/db"
	"dballz/internal/model"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(database db.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (r *UserRepository) ListUsers() ([]model.User, error) {
	rows, err := r.db.Conn().Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	return users, nil
}
