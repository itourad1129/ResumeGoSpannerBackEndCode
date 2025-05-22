package controller

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"net/http"
	"pjdrc/domain"
	"strconv"
)

type MasterDataVersionController struct {
	SpannerClient            *spanner.Client
	MasterDataVersionUsecase domain.MasterDataVersionUsecase
}

func (mdc *MasterDataVersionController) GetMasterDataVersion(c *gin.Context) {

	masterDataVersions, err := mdc.MasterDataVersionUsecase.GetMasterDataVersion(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var masterDataVersionResponse []domain.GetMasterDataVersionResponse
	for _, masterDataVersion := range masterDataVersions {
		response := domain.GetMasterDataVersionResponse{
			MasterDataID: strconv.FormatInt(masterDataVersion.MasterDataID, 10),
			Version:      strconv.FormatInt(masterDataVersion.Version, 10),
			ChunkID:      strconv.FormatInt(masterDataVersion.ChunkID, 10),
		}
		masterDataVersionResponse = append(masterDataVersionResponse, response)
	}

	c.JSON(http.StatusOK, masterDataVersionResponse)
}
