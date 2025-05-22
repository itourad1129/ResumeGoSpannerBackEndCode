package route

import (
	"cloud.google.com/go/spanner"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pjdrc/api/controller"
	"pjdrc/repository"
	"pjdrc/usecase"
	"time"
)

func UserLoginRouter(timeout time.Duration, handleJwtMiddleware *jwt.GinJWTMiddleware, group *gin.RouterGroup, spannerClient *spanner.Client) {

	ulr := repository.NewUserLoginRepository(spannerClient, "t_user_login")
	uar := repository.NewUserAreaRepository(spannerClient, "t_user_area")
	ulc := controller.UserLoginController{
		SpannerClient:    spannerClient,
		UserLoginUsecase: usecase.NewUserLoginUsecase(ulr, uar, timeout),
		JwtMiddleware:    handleJwtMiddleware,
	}
	group.POST("/login", ulc.UserLoginHandler)
}
