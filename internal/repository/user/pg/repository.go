package pg

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/user/pg/converter"
	modelRepo "github.com/greenblat17/auth/internal/repository/user/pg/model"
	"github.com/greenblat17/platform-common/pkg/db"
	"github.com/jackc/pgx/v4"
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

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	sqb := sq.Insert(userTable).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(info.Name, info.Email, info.Password, info.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "UserRepository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().ScanOneContext(ctx, &userID, q, args...)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) Update(ctx context.Context, user *model.User) error {
	sqb := sq.Update(userTable).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID}).
		PlaceholderFormat(sq.Dollar)

	if !user.Info.IsEmptyName() {
		sqb = sqb.Set(nameColumn, user.Info.Name)
	}
	if !user.Info.IsEmptyEmail() {
		sqb = sqb.Set(emailColumn, user.Info.Email)
	}
	if !user.Info.IsEmptyRole() {
		sqb = sqb.Set(roleColumn, user.Info.Role)
	}

	query, args, err := sqb.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UserRepository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(ctx context.Context, filter *model.UserFilter) (*model.User, error) {
	sqb := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(userTable).
		PlaceholderFormat(sq.Dollar)

	sqb = applyFilter(sqb, filter)

	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "UserRepository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrUserNotFound
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

	q := db.Query{
		Name:     "UserRepository.Delete",
		QueryRaw: query,
	}

	cmdTag, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := cmdTag.RowsAffected()
	if rowsAffected == 0 {
		return repository.ErrUserNotFound
	}

	return nil
}

func applyFilter(sqb sq.SelectBuilder, filter *model.UserFilter) sq.SelectBuilder {
	if filter == nil {
		return sqb
	}

	if len(filter.Name) != 0 {
		sqb = sqb.Where(sq.Eq{nameColumn: filter.Name})
	}
	if filter.ID != 0 {
		sqb = sqb.Where(sq.Eq{idColumn: filter.ID})
	}

	return sqb
}
