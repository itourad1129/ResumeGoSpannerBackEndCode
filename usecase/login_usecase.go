package usecase

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain"
	"pjdrc/domain/user"
	"time"
)

type userLoginUsecase struct {
	userLoginRepository user.UserLoginRepository
	userAreaRepository  user.UserAreaRepository
	contextTimeout      time.Duration
}

func NewUserLoginUsecase(userLoginRepository user.UserLoginRepository, userAreaRepository user.UserAreaRepository, timeout time.Duration) domain.UserLoginUsecase {
	return &userLoginUsecase{
		userLoginRepository: userLoginRepository,
		userAreaRepository:  userAreaRepository,
		contextTimeout:      timeout,
	}
}

func (u userLoginUsecase) InsertOrUpdate(c context.Context, tx *spanner.ReadWriteTransaction, userID int64) (user.UserLogin, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userLoginRepository.InsertOrUpdate(ctx, tx, userID)
}

func (u userLoginUsecase) GetUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID string) (user.UserArea, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userAreaRepository.GetUserArea(ctx, tx, userID)
}
