package converter

import (
	"github.com/greenblat17/auth/internal/model"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromService converts user model to proto
func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserInfoFromService converts user info model to proto
func ToUserInfoFromService(userInfo model.UserInfo) *desc.UserInfo {
	var role desc.Role
	switch userInfo.Role {
	case model.RoleAdmin:
		role = desc.Role_ADMIN
	case model.RoleUser:
		role = desc.Role_USER
	default:
		role = desc.Role_UNKNOWN
	}

	return &desc.UserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Role:     role,
		Password: userInfo.Password,
	}
}

// ToUserFromAPI converts proto to user model
func ToUserFromAPI(user *desc.User) *model.User {
	updatedAt := user.GetUpdatedAt().AsTime()

	return &model.User{
		ID:        user.GetId(),
		Info:      ToUserInfoFromAPI(user.Info),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: &updatedAt,
	}
}

// ToUserInfoFromAPI converts proto to user info model
func ToUserInfoFromAPI(userInfo *desc.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:     userInfo.GetName(),
		Email:    userInfo.GetEmail(),
		Role:     userInfo.GetRole().String(),
		Password: userInfo.GetPassword(),
	}
}
