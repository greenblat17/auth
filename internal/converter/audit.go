package converter

import "github.com/greenblat17/auth/internal/model"

// ToAuditFromEntity convert entity to audit model
func ToAuditFromEntity(entity model.EntityType, action string) *model.Audit {
	return &model.Audit{
		Entity: entity,
		Action: action,
	}
}
