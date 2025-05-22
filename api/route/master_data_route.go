package route

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"pjdrc/api/controller"
	"pjdrc/repository"
	"pjdrc/usecase"
	"time"
)

func MasterDataRouter(timeout time.Duration, group *gin.RouterGroup, spannerClient *spanner.Client) {
	cvr := repository.NewMasterDataVersionRepository(spannerClient, "m_master_data_version")
	cr := controller.MasterDataVersionController{
		SpannerClient:            spannerClient,
		MasterDataVersionUsecase: usecase.NewMasterDataUsecase(cvr, timeout),
	}
	group.GET("/getMasterDataVersion", cr.GetMasterDataVersion)
}
