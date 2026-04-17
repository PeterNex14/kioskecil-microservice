package service

import (
	"context"

	"github.com/google/uuid"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"
	"github.com/PeterNex14/kioskecil-microservice/user-service/internal/repository"
)

// UserService defines the contract for business logic regarding users
type UserService interface {
	RegisterUser(ctx context.Context, arg db_users_gen.CreateUserParams) (db_users_gen.User, error)
	GetByEmail(ctx context.Context, email string) (db_users_gen.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (db_users_gen.User, error)
}

// userService is the implementation of UserService
type userService struct {
	repo repository.UserRepository
}

// NewUserService returns a new instance of UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) RegisterUser(ctx context.Context, arg db_users_gen.CreateUserParams) (db_users_gen.User, error) {
	// TODO: Implement password hashing logic here before calling the repository
	return s.repo.CreateUser(ctx, arg)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (db_users_gen.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (db_users_gen.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
