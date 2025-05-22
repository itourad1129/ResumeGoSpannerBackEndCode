package domain

import (
	"context"
	"pjdrc/domain/master"
)

type ChunkUsecase interface {
	GetChunkVersion(c context.Context, platformType int64) (master.ChunkVersion, error)
}

type GetChunkVersionRequest struct {
	PlatformType int64 `form:"platformType" binding:"required"`
}

type GetChunkVersionResponse struct {
	VersionID      string `json:"versionID"`
	PlatformType   string `json:"platformType"`
	DeploymentName string `json:"deploymentName"`
	ContentBuildID string `json:"contentBuildID"`
}
