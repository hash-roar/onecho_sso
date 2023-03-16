package userservice

import (
	"context"
	"onecho_sso_backend/models"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/logging"
	"onecho_sso_backend/utils"

	"go.uber.org/zap"
)

func RegisterByName(ctx context.Context, nickName string, password string) int {
	if nickName == "" || password == "" {
		return enums.INVALID_PARAMS
	}
	user, err := models.GetUserByNickName(nickName)
	if err != nil {
		logging.Ctx(ctx).Error("error in find user by name", zap.Error(err))
		return enums.SERVER_ERROR
	}
	// if user exists
	if user != nil {
		return enums.USER_NAME_DUPLICATED
	}
	// add user
	encryptPass, err := utils.EncryptPassword(password)
	if err != nil {
		logging.Ctx(ctx).Error("error in encrypt password", zap.Error(err))
		return enums.SERVER_ERROR
	}
	err = models.AddUser(&models.User{
		NickName: nickName,
		Password: encryptPass,
	})

	if err != nil {
		logging.Ctx(ctx).Error("error in add user", zap.Error(err))
		return enums.SERVER_ERROR
	}

	return enums.SUCCESS
}

func RegisterByEmail(ctx context.Context, email string, password string) int {
	if email == "" || password == "" {
		return enums.INVALID_PARAMS
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		logging.Ctx(ctx).Error("error in find user by name", zap.Error(err))
		return enums.SERVER_ERROR
	}
	// if user exists
	if user != nil {
		return enums.USER_EXIST
	}
	encryptPass, err := utils.EncryptPassword(password)
	if err != nil {
		logging.Ctx(ctx).Error("error in encrypt password", zap.Error(err))
		return enums.SERVER_ERROR
	}
	err = models.AddUser(&models.User{
		Email:    email,
		Password: encryptPass,
		Name:     email,
	})

	if err != nil {
		logging.Ctx(ctx).Error("error in add user", zap.Error(err))
		return enums.SERVER_ERROR
	}

	return enums.SUCCESS

}
