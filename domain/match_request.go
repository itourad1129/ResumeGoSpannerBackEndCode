package domain

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain/master"
	"pjdrc/domain/user"
)

type MatchRequestUsecase interface {
	GetArea(c context.Context, areaID int64) (master.Area, error)
	GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID string) (user.UserArea, error)
}

type MatchRequestResponse struct {
	UserID  string `json:"userID"`
	AreaID  string `json:"areaID"`
	LevelID string `json:"levelID"`
}
