// Package router provides routing functionality for the application.
package router

import (
	"github.com/aleodoni/voting-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(jwtMiddleware *middleware.JWTMiddleware) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")

	registerHealthRoutes(api)
	registerProtectedRoutes(api, jwtMiddleware)

	return r
}
