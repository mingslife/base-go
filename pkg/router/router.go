package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"base-go/pkg/conf"
	"base-go/pkg/controllers"
	"base-go/pkg/middleware"
)

func NewRouter(cfg *conf.Config) *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api/v1")
	apiRouter.Use(middleware.NewCorsMiddleware())
	apiRouter.Use(middleware.NewJwtMiddleware(&middleware.JwtMiddlewareConfig{
		Realm:      cfg.Name,
		Key:        []byte(cfg.JwtKey),
		Timeout:    time.Duration(cfg.JwtTimeout) * time.Hour,
		MaxRefresh: time.Duration(cfg.JwtMaxRefresh) * time.Hour,
		ExcludePaths: []string{
			"/api/v1/auth/login",
		},
	}))
	controllers.NewAuthController(apiRouter)
	controllers.NewUserController(apiRouter)

	return router
}
