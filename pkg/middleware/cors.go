package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware() gin.HandlerFunc {
	return cors.Default()
}

type corsNS struct{}

var CORS corsNS
