package domain

import (
	"context"
	"pjdrc/domain/master"
)

type MasterDataVersionUsecase interface {
	GetMasterDataVersion(c context.Context) ([]master.MasterDataVersion, error)
}

type GetMasterDataVersionResponse struct {
	MasterDataID string `json:"masterDataid"`
	Version      string `json:"version"`
	ChunkID      string `json:"chunkID"`
}
