package grpcApi

import (
	"context"
	"github.com/hawkkiller/wtc-account-service-api/api"
	internal "github.com/hawkkiller/wtc-account/internal/database"
	"github.com/hawkkiller/wtc-account/internal/model"
)

func (AccountService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	id := req.GetId()
	userDB := new(model.UserProfile)
	if res := internal.DB.Where("id = ?", id).First(&userDB); res.Error != nil {
		return nil, res.Error
	}

	return &api.GetUserResponse{
		Username: userDB.Username,
		Email:    userDB.Email,
		Sex:      userDB.Sex,
	}, nil
}
