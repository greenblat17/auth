package access

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/access/converter"
	modelRepo "github.com/greenblat17/auth/internal/repository/access/model"
	"github.com/greenblat17/platform-common/pkg/db"
	"github.com/jackc/pgx/v4"
)

const (
	accessRuleTable = "access_rules"

	roleColumn     = "role"
	endpointColumn = "endpoint"
)

type accessRuleRepo struct {
	db db.Client
}

// NewRepository creates a new access rule repository
func NewRepository(db db.Client) *accessRuleRepo {
	return &accessRuleRepo{db: db}
}

func (r *accessRuleRepo) GetByEndpoint(ctx context.Context, endpoint string) (*model.AccessRule, error) {
	sqb := sq.Select(endpointColumn, roleColumn).
		From(accessRuleTable).
		Where(sq.Eq{endpointColumn: endpointColumn}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := sqb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "AccessRuleRepo.GetByEndpoint",
		QueryRaw: query,
	}

	var rules []*modelRepo.AccessRule
	err = r.db.DB().ScanAllContext(ctx, &rules, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrRuleNotFound
		}

		return nil, err
	}

	return converter.ToAccessRuleFromRepo(endpoint, rules), nil
}
