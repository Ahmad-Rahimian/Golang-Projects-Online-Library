package user

import (
	"context"
	"database/sql"
)

// Repository interface for user repository
type Repository interface {
	CreateUser(ctx context.Context, phone string) error
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}

// repository struct for user repository
type repository struct {
	db *sql.DB
}

// new repository for user repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// when create new user, default role is user
func (r *repository) CreateUser(ctx context.Context, phone string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (phone, role) VALUES ($1, $2)", phone, "user")
	return err
}

// get user by phone from database and return user
func (r *repository) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, phone, role FROM users WHERE phone=$1", phone)
	var u User
	if err := row.Scan(&u.ID, &u.Phone, &u.Role); err != nil {
		return nil, err
	}
	return &u, nil
}
