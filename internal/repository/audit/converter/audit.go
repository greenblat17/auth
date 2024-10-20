package converter

import (
	"github.com/greenblat17/auth/internal/model"
	modelRepo "github.com/greenblat17/auth/internal/repository/audit/model"
)

// ToAuditFromService convert audit model from service to repo entity
func ToAuditFromService(audit *model.Audit) *modelRepo.Audit {
	return &modelRepo.Audit{
		Entity:    string(audit.Entity),
		Action:    audit.Action,
		CreatedAt: audit.CreatedAt,
	}
}

// ToAuditFromRepo convert audit model from repository to service entity
func ToAuditFromRepo(audit *modelRepo.Audit) *model.Audit {
	return &model.Audit{
		Entity:    model.EntityType(audit.Entity),
		Action:    audit.Action,
		CreatedAt: audit.CreatedAt,
	}
}
