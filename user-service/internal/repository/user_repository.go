package repository

import (
	"context"

	"github.com/google/uuid"
	db_users_gen "github.com/PeterNex14/kioskecil-microservice/user-service/db/sqlc"
)

// UserRepository defines the contract for user data access
type UserRepository interface {
	CreateUser(ctx context.Context, arg db_users_gen.CreateUserParams) (db_users_gen.User, error)
	GetUserByEmail(ctx context.Context, email string) (db_users_gen.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (db_users_gen.User, error)
}

// userRepository is the SQLC implementation of UserRepository
type userRepository struct {
	queries *db_users_gen.Queries
}

// NewUserRepository returns a new instance of UserRepository
func NewUserRepository(queries *db_users_gen.Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, arg db_users_gen.CreateUserParams) (db_users_gen.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (db_users_gen.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (db_users_gen.User, error) {
	return r.queries.GetUserByID(ctx, id)
}
