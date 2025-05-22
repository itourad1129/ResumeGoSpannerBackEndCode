package domain

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain/user"
)

type UserRegisterRequest struct {
	Name string `form:"name" binding:"required"`
}

type UserRegisterResponse struct {
	UserID       string `json:"userID"`
	TransferCode string `json:"transferCode"`
}

type UserRegisterUsecase interface {
	CreateUserInfo(c context.Context, tx *spanner.ReadWriteTransaction, userName string) (int64, error)
	CreateUserTransfer(c context.Context, tx *spanner.ReadWriteTransaction, userID int64) (string, error)
	CreateUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID int64) error
	GetUserByUserName(c context.Context, tx *spanner.ReadWriteTransaction, name string) (user.UserInfo, error)
}
