package controller

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"net/http"
	"pjdrc/domain"
	"strconv"
)

type ChunkController struct {
	SpannerClient *spanner.Client
	ChunkUsecase  domain.ChunkUsecase
}

func (cc *ChunkController) GetChunkVersion(c *gin.Context) {

	var request domain.GetChunkVersionRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	chunkVersion, err := cc.ChunkUsecase.GetChunkVersion(c, request.PlatformType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	chunkVersionResponse := domain.GetChunkVersionResponse{
		VersionID:      strconv.FormatInt(chunkVersion.VersionID, 10),
		PlatformType:   strconv.FormatInt(chunkVersion.PlatformType, 10),
		DeploymentName: chunkVersion.DeploymentName,
		ContentBuildID: chunkVersion.ContentBuildID,
	}

	c.JSON(http.StatusOK, chunkVersionResponse)
}
