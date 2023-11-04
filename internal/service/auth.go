package service

import "github.com/strwys/fms/internal/repository"

type AuthService interface {
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}
