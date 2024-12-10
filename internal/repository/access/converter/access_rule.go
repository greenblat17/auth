package converter

import (
	"github.com/greenblat17/auth/internal/model"
	modelRepo "github.com/greenblat17/auth/internal/repository/access/model"
)

// ToAccessRuleFromRepo converts repo entity to access rule
func ToAccessRuleFromRepo(endpoint string, rules []*modelRepo.AccessRule) *model.AccessRule {
	accessRule := &model.AccessRule{
		Endpoint: endpoint,
		Role:     make(map[string]struct{}, len(rules)),
	}

	for _, rule := range rules {
		accessRule.Role[rule.Role] = struct{}{}
	}

	return accessRule
}
