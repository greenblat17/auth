package audit

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/greenblat17/auth/internal/client/db"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
)

const (
	auditTable = "audit"

	idColumn        = "id"
	entityColumn    = "entity"
	actionColumn    = "action"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository returnrs audit repository
func NewRepository(db db.Client) repository.AuditRepository {
	return &repo{db: db}
}

func (r *repo) Save(ctx context.Context, audit *model.Audit) error {
	sqb := sq.Insert(auditTable).
		Columns(entityColumn, actionColumn, createdAtColumn).
		Values(audit.Entity, audit.Action, time.Now()).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "AuditRepository.Save",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
