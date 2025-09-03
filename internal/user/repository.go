package user

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
	u := &User{}
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}
