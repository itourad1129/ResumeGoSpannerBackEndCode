package usecase

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain"
	"pjdrc/domain/master"
	"pjdrc/domain/user"
	"time"
)

type matchRequestUsecase struct {
	areaRepository     master.AreaRepository
	userAreaRepository user.UserAreaRepository
	contextTimeout     time.Duration
}

func NewMatchRequestUsecase(areaRepository master.AreaRepository, userAreaRepository user.UserAreaRepository, timeout time.Duration) domain.MatchRequestUsecase {
	return &matchRequestUsecase{
		areaRepository:     areaRepository,
		userAreaRepository: userAreaRepository,
		contextTimeout:     timeout,
	}
}

func (u matchRequestUsecase) GetArea(c context.Context, areaID int64) (master.Area, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.areaRepository.GetArea(ctx, areaID)
}

func (u matchRequestUsecase) GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID string) (user.UserArea, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userAreaRepository.GetUserArea(ctx, tx, userID)
}
