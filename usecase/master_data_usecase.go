package usecase

import (
	"context"
	"pjdrc/domain"
	"pjdrc/domain/master"
	"time"
)

type masterDataVersionUsecase struct {
	masterDataVersionRepository master.MasterDataVersionRepository
	contextTimeout              time.Duration
}

func NewMasterDataUsecase(masterDataVersionRepository master.MasterDataVersionRepository, timeout time.Duration) domain.MasterDataVersionUsecase {
	return &masterDataVersionUsecase{
		masterDataVersionRepository: masterDataVersionRepository,
		contextTimeout:              timeout,
	}
}

func (u masterDataVersionUsecase) GetMasterDataVersion(c context.Context) ([]master.MasterDataVersion, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.masterDataVersionRepository.GetMasterDataVersion(ctx)
}
