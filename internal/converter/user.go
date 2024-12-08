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
		Name:  userInfo.Name,
		Email: userInfo.Email,
		Role:  role,
	}
}

// ToUserInfoFromCreateAPI convert create request proto to user info model
func ToUserInfoFromCreateAPI(req *desc.CreateRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole().String(),
	}
}

// ToUserFromUpdateAPI convert update request proto to user model
func ToUserFromUpdateAPI(req *desc.UpdateRequest) *model.User {
	return &model.User{
		ID: req.GetId(),
		Info: model.UserInfo{
			Name:  req.GetName().GetValue(),
			Email: req.GetEmail().GetValue(),
			Role:  req.GetRole().String(),
		},
	}
}
