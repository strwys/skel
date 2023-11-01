package repository

import "database/sql"

type AuthRepository interface {
}

type authRepository struct {
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{}
}
