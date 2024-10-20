package converter

import (
	"time"

	"github.com/greenblat17/auth/internal/model"
	modelRepo "github.com/greenblat17/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.Info) model.UserInfo {
	return model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Role:     info.Role,
		Password: info.Password,
	}
}
