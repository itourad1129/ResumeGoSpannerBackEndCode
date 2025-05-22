package master

import (
	"context"
)

const (
	AreaID  = "AreaID"
	LevelID = "LevelID"
)

type Area struct {
	AreaID  int64 `spanner:"area_id"`
	LevelID int64 `spanner:"level_id"`
}

type AreaRepository interface {
	GetArea(c context.Context, areaID int64) (Area, error)
}
