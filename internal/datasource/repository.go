package datasource

import (
	"MyProgy/infrastructure/database"
	"MyProgy/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	CreateUser(ctx context.Context, u domain.User) (int, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, id int, u domain.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByName(ctx context.Context, name, password string) (domain.User, error)
}

type Repository struct {
	Client database.Client
}

func NewRepository(client database.Client) Storage {
	return &Repository{
		Client: client,
	}
}

func (r *Repository) CreateUser(ctx context.Context, u domain.User) (id int, err error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = r.Client.QueryRow(ctx, query, u.Name, u.Email, passwordHash).Scan(&id)
	return id, err
}

func (r *Repository) GetUsers(ctx context.Context) ([]domain.User, error) {
	guery := `SELECT id, name, email FROM users`
	rows, err := r.Client.Query(ctx, guery)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return users, nil
}

func (r *Repository) GetUserById(ctx context.Context, id int) (domain.User, error) {
	guery := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.Client.QueryRow(ctx, guery, id)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByName(ctx context.Context, name, password string) (domain.User, error) {
	guery := `SELECT id, name, email, password FROM users WHERE name = $1`
	row := r.Client.QueryRow(ctx, guery, name)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash); err != nil {
		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}
	err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return domain.User{}, fmt.Errorf("%s, %s", user.PasswordHash, password)
	}
	return user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id int, u domain.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id`
	err := r.Client.QueryRow(ctx, query, u.Name, u.Email, id).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("user with id %d not found", id)
		}
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1 RETURNING id`
	err := r.Client.QueryRow(ctx, query, id).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("user with id %d not found", id)
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
