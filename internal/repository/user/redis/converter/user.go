package converter

import (
	"time"

	"github.com/greenblat17/auth/internal/model"
	repoModel "github.com/greenblat17/auth/internal/repository/user/redis/model"
)

// ToUserFromService convert user service entity to redis entity
func ToUserFromService(user *model.User) *repoModel.User {
	var updatedAt int64
	if user.UpdatedAt != nil {
		updatedAt = user.UpdatedAt.UnixNano()
	}

	return &repoModel.User{
		ID:          user.ID,
		Name:        user.Info.Name,
		Email:       user.Info.Email,
		Role:        user.Info.Role,
		CreatedAtNs: user.CreatedAt.UnixNano(),
		UpdatedAtNs: &updatedAt,
	}
}

// ToUserFromRepo convert user redis entity to service entity
func ToUserFromRepo(user *repoModel.User) *model.User {
	var updatedAt time.Time
	if user.UpdatedAtNs != nil {
		updatedAt = time.Unix(0, *user.UpdatedAtNs)
	}

	return &model.User{
		ID: user.ID,
		Info: model.UserInfo{
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		CreatedAt: time.Unix(0, user.CreatedAtNs),
		UpdatedAt: &updatedAt,
	}
}
