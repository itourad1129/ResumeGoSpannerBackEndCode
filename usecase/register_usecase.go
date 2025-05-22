package usecase

import (
	"cloud.google.com/go/spanner"
	"context"
	"pjdrc/domain"
	"pjdrc/domain/user"
	"time"
)

type userRegisterUsecase struct {
	userInfoRepository     user.UserInfoRepository
	userTransferRepository user.UserTransferRepository
	userAreaRepository     user.UserAreaRepository
	contextTimeout         time.Duration
}

func NewUserRegisterUsecase(userInfoRepository user.UserInfoRepository, userTransferRepository user.UserTransferRepository, userAreaRepository user.UserAreaRepository, timeout time.Duration) domain.UserRegisterUsecase {
	return &userRegisterUsecase{
		userInfoRepository:     userInfoRepository,
		userTransferRepository: userTransferRepository,
		userAreaRepository:     userAreaRepository,
		contextTimeout:         timeout,
	}
}

func (u *userRegisterUsecase) CreateUserInfo(c context.Context, tx *spanner.ReadWriteTransaction, userName string) (int64, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userInfoRepository.Create(ctx, tx, userName)
}

func (u *userRegisterUsecase) CreateUserTransfer(c context.Context, tx *spanner.ReadWriteTransaction, userID int64) (string, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userTransferRepository.Create(ctx, tx, userID)
}

func (u *userRegisterUsecase) CreateUserArea(c context.Context, tx *spanner.ReadWriteTransaction, userID int64) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userAreaRepository.Create(ctx, tx, userID)
}

func (u *userRegisterUsecase) GetUserByUserName(c context.Context, tx *spanner.ReadWriteTransaction, email string) (user.UserInfo, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userInfoRepository.GetUserName(ctx, tx, email)
}

//
//func (su *userRegisterUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
//	return tokenutil.CreateAccessToken(user, secret, expiry)
//}
//
//func (su *userRegisterUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
//	return tokenutil.CreateRefreshToken(user, secret, expiry)
//}
