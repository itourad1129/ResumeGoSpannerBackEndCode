package route

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"pjdrc/api/controller"
	"pjdrc/repository"
	"pjdrc/usecase"
	"time"
)

func MatchRequestRouter(timeout time.Duration, group *gin.RouterGroup, spannerClient *spanner.Client) {

	uar := repository.NewUserAreaRepository(spannerClient, "t_user_area")
	ar := repository.NewAreaRepository(spannerClient, "m_area")
	mrc := controller.MatchRequestController{
		SpannerClient:       spannerClient,
		MatchRequestUsecase: usecase.NewMatchRequestUsecase(ar, uar, timeout),
	}
	group.GET("/matchRequest", mrc.MatchRequest)
}
