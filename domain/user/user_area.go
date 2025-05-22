package user

import (
	"cloud.google.com/go/spanner"
	"context"
)

const (
	AreaID = "AreaID"
	IsStay = "IsStay"
)

type UserArea struct {
	UserID int64 `spanner:"user_id"`
	AreaID int64 `spanner:"area_id"`
	IsStay bool  `spanner:"is_stay"`
}

type UserAreaRepository interface {
	Create(c context.Context, tx *spanner.ReadWriteTransaction, userId int64) error
	GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userId string) (UserArea, error)
}
