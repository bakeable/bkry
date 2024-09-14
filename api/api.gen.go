package api

import (
	"github.com/bakeable/bkry/api/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for different models
func InitRoutes(r *gin.RouterGroup) {
	r.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))
	r.Use(middleware.RecoveryWithWriter())

	// Initialize custom routes
	InitCustomRoutes(r)
}
