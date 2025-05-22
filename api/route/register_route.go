package route

import (
	"cloud.google.com/go/spanner"
	"github.com/gin-gonic/gin"
	"pjdrc/api/controller"
	"pjdrc/repository"
	"pjdrc/usecase"
	"time"
)

func UserRegisterRouter(timeout time.Duration, group *gin.RouterGroup, spannerClient *spanner.Client) {
	uir := repository.NewUserInfoRepository(spannerClient, "t_user_info")
	utr := repository.NewUserTransferRepository(spannerClient, "t_user_transfer")
	uar := repository.NewUserAreaRepository(spannerClient, "t_user_area")
	urc := controller.UserRegisterController{
		SpannerClient:       spannerClient,
		UserRegisterUsecase: usecase.NewUserRegisterUsecase(uir, utr, uar, timeout),
	}
	group.POST("/userRegister", urc.UserRegister)
}
