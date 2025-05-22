package controller

import (
	"cloud.google.com/go/spanner"
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"pjdrc/domain"
	"strconv"
)

type MatchRequestController struct {
	SpannerClient       *spanner.Client
	MatchRequestUsecase domain.MatchRequestUsecase
}

func (mrc *MatchRequestController) MatchRequest(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	ClaimUserID := claims["userID"]

	// JSONパース後などは float64 で保持されていることが多い
	userIDFloat, ok := ClaimUserID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "match request parse error"})
		return
	}
	userID := int64(userIDFloat)

	_, err := mrc.SpannerClient.ReadWriteTransaction(c, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {

		userArea, err := mrc.MatchRequestUsecase.GetUserArea(c, tx, strconv.FormatInt(userID, 10))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return err
		}

		area, err := mrc.MatchRequestUsecase.GetArea(c, userArea.AreaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return err
		}

		matchRequestResponse := domain.MatchRequestResponse{
			UserID:  strconv.FormatInt(userID, 10),
			AreaID:  strconv.FormatInt(area.AreaID, 10),
			LevelID: strconv.FormatInt(area.LevelID, 10),
		}
		c.Set("matchRequestResponse", matchRequestResponse)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if matchRequestResponse, exists := c.Get("matchRequestResponse"); exists {
		c.JSON(http.StatusOK, matchRequestResponse)
	}
}
