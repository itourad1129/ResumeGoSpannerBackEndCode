package domain

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain/user"
)

type UserLoginRequest struct {
	UserID       int64  `form:"userID" json:"userID" binding:"required"`
	TransferCode string `form:"transferCode" json:"transferCode" binding:"required"`
}

type UserLoginResponse struct {
	UserID         string `json:"userID"`
	TotalLoginDays string `json:"totalLoginDays"`
	AreaID         string `json:"areaID"`
}

type UserLoginUsecase interface {
	InsertOrUpdate(c context.Context, tx *spanner.ReadWriteTransaction, userId int64) (user.UserLogin, error)
	GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID string) (user.UserArea, error)
}
