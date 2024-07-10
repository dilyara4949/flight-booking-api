package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db: db}
}

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrNothingChanged = errors.New("nothing changed")
)

const (
	createUser  = "insert into users (id, email, password, phone, role) values ($1, $2, $3, $4, 'user') returning created_at, updated_at"
	getUser     = "select id, email, phone, created_at, updated_at from users where id = $1;"
	updateUser  = "update users set email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP where id = $1;"
	deleteUsers = "delete from users where id = $1"
	getAllUsers = "select id, email, phone, created_at, updated_at from users limit $1 offset $2;"
)

func (repo *UserRepository) Create(ctx context.Context, user *domain.User, password string) error {
	if err := repo.db.QueryRowContext(ctx, createUser, user.ID, user.Email, password, user.Phone).Scan(&user.CreatedAt, &user.UpdatedAt); err != nil {
		return fmt.Errorf("create user error: %v", err)
	}
	return nil
}

func (repo *UserRepository) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row := repo.db.QueryRowContext(ctx, getUser, id)

	user := domain.User{}

	err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user error: %v", err)
	}

	return &user, nil
}

func (repo *UserRepository) Update(ctx context.Context, user domain.User) error {
	res, err := repo.db.ExecContext(ctx, updateUser, user.ID, user.Email, user.Phone)
	if err != nil {
		return fmt.Errorf("update user error: %v", err)
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (repo *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res, err := repo.db.ExecContext(ctx, deleteUsers, id)
	if err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (repo *UserRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.User, error) {
	offset := (page - 1) * pageSize

	rows, err := repo.db.QueryContext(ctx, getAllUsers, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("get all users error: %v", err)
	}
	defer rows.Close()

	users := make([]domain.User, 0)

	for rows.Next() {
		user := domain.User{}

		err = rows.Scan(&user.ID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
