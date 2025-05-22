package route

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"pjdrc/api/controller"
	"pjdrc/repository"
	"pjdrc/usecase"
	"time"
)

func ChunkRouter(timeout time.Duration, group *gin.RouterGroup, spannerClient *spanner.Client) {
	cvr := repository.NewChunkVersionRepository(spannerClient, "m_chunk_version")
	cr := controller.ChunkController{
		SpannerClient: spannerClient,
		ChunkUsecase:  usecase.NewChunkUsecase(cvr, timeout),
	}
	group.POST("/getChunkVersion", cr.GetChunkVersion)
}
