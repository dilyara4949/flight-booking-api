package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dilyara4949/flight-booking-api/internal/domain"
	"github.com/google/uuid"
)

type userRepository struct {
	db             *sql.DB
	contextTimeout time.Duration
}

func NewUserRepository(db *sql.DB, timeout time.Duration) domain.UserRepository {
	return &userRepository{db: db, contextTimeout: timeout}
}

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrNothingChanged = errors.New("nothing changed")
)

const (
	createUser  = "insert into users (id, email, password, phone, role_id) values ($1, $2, $3, $4, $5) returnig created_at, updated_at"
	getUser     = "select id, email, phone, created_at, updated_at, role_id from users where id = $1;"
	updateUser  = "update users set email = $2, phone = $3, updated_at = CURRENT_TIMESTAMP where id = $1;"
	deleteUsers = "delete from users where id = $1"
	getAllUsers = "select id, email, phone, role_id, created_at, updated_at from users limit $1 offset $2;"
)

func (repo *userRepository) Create(ctx context.Context, user domain.User, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout*time.Second)
	defer cancel()

	user.ID = uuid.New().String()

	if err := repo.db.QueryRowContext(ctx, createUser, user.ID, user.Email, password, user.Phone).Scan(&user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout*time.Second)
	defer cancel()

	row := repo.db.QueryRowContext(ctx, getUser, id)

	user := domain.User{}

	err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) Update(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout*time.Second)
	defer cancel()

	res, err := repo.db.ExecContext(ctx, updateUser, user.ID, user.Email, user.Phone)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (repo *userRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout*time.Second)
	defer cancel()

	res, err := repo.db.ExecContext(ctx, deleteUsers, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (repo *userRepository) GetAll(ctx context.Context, page, pageSize int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout*time.Second)
	defer cancel()

	offset := (page - 1) * pageSize

	rows, err := repo.db.QueryContext(ctx, getAllUsers, pageSize, offset)
	if err != nil {
		return nil, err
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
