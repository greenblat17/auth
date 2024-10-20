package user

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/user/converter"
	modelRepo "github.com/greenblat17/auth/internal/repository/user/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	userTable = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

var (
	// ErrNotFound is returned when a user is not found.
	ErrNotFound = errors.New("user not found")
)

type repo struct {
	db *pgxpool.Pool
}

// NewRepository creates a new user repository.
func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	sqb := sq.Insert(userTable).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Info.Name, user.Info.Email, user.Info.Password, user.Info.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) Update(ctx context.Context, user *model.User) error {
	sqb := sq.Update(userTable).
		Set(nameColumn, user.Info.Name).
		Set(emailColumn, user.Info.Email).
		Set(roleColumn, user.Info.Role).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	sqb := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(userTable).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, err
	}

	var user modelRepo.User
	err = r.db.QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Password, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	sqb := sq.Delete(userTable).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return err
	}

	cmdTag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected := cmdTag.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
