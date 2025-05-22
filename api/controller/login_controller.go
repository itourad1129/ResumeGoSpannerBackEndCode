package controller

import (
	"cloud.google.com/go/spanner"
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"pjdrc/api/middleware"
	"pjdrc/domain"
	"strconv"
)

type UserLoginController struct {
	SpannerClient    *spanner.Client
	UserLoginUsecase domain.UserLoginUsecase
	JwtMiddleware    *jwt.GinJWTMiddleware
}

func (ulc *UserLoginController) UserLoginInternal(c *gin.Context, userID int64) (domain.UserLoginResponse, error) {
	var result domain.UserLoginResponse
	_, err := ulc.SpannerClient.ReadWriteTransaction(c, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		userLogin, err := ulc.UserLoginUsecase.InsertOrUpdate(c, tx, userID)
		userArea, err := ulc.UserLoginUsecase.GetUserArea(ctx, tx, strconv.FormatInt(userID, 10))
		if err != nil {
			return err
		}

		result = domain.UserLoginResponse{
			UserID:         strconv.FormatInt(userID, 10),
			TotalLoginDays: strconv.FormatInt(userLogin.TotalLoginDays, 10),
			AreaID:         strconv.FormatInt(userArea.AreaID, 10),
		}
		return nil
	})
	return result, err
}

func (ulc *UserLoginController) UserLoginHandler(c *gin.Context) {
	// 認証処理
	claims, err := ulc.JwtMiddleware.Authenticator(c)
	if err != nil {
		ulc.JwtMiddleware.Unauthorized(c, http.StatusUnauthorized, err.Error())
		return
	}

	// トークン生成
	token, expire, err := ulc.JwtMiddleware.TokenGenerator(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "create token error" + err.Error()})
		return
	}

	// DB処理
	userID := middleware.LoginUserID
	userLogin, err := ulc.UserLoginInternal(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// 結果まとめて返却
	c.JSON(http.StatusOK, gin.H{
		"code":           http.StatusOK,
		"expire":         expire,
		"token":          token,
		"userID":         userLogin.UserID,
		"totalLoginDays": userLogin.TotalLoginDays,
		"areaID":         userLogin.AreaID,
	})
}
