package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pjdrc/api/middleware"
	"pjdrc/api/route"
	"pjdrc/database"
	"pjdrc/domain"
	mytime "pjdrc/domain/time"
	"time"
)

//const ServerEnv = ""

func main() {

	// デバッグ用時間変更
	var setDuration time.Duration = 0
	setDuration = -365 * 24 * time.Hour      // 1年前
	setDuration = setDuration + 24*time.Hour //1日後
	mytime.SetOffset(setDuration)
	fmt.Println("現在時刻:", mytime.Now())

	domain.FlagInit()
	timeout := 5 * time.Second

	client, err := database.NewSpannerClient()
	if err != nil {
		log.Fatalf("Failed to create Spanner client: %v", err)
	}
	defer client.Close()

	r := gin.Default()
	authMiddleware, err := middleware.NewJwtMiddleware(client)
	if err != nil {
		log.Fatalf("Failed to create Jwt middleware: %v", err)
	}

	r.Use(middleware.HandlerJwtMiddleWare(authMiddleware))
	route.Setup(r, authMiddleware, timeout, client)
	r.Run(":8080")
}
