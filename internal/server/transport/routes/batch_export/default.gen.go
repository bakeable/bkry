package batch_export_routes

import (
	batch_export_controllers "github.com/bakeable/bkry/internal/server/transport/controllers/batch_export"
	"github.com/bakeable/bkry/internal/server/transport/middleware"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func getControllers() batch_export_controllers.IControllers {
	return batch_export_controllers.NewControllers()
}

func InitRoutes(r *gin.RouterGroup) {
	// Get controllers
	controllers := getControllers()

	// Create a new route group for the BatchExport entity
	group := r.Group("/batch_export")

	// Add authentication middleware for all endpoints
	group.Use(middleware.AuthenticatedRoutes())
	
	// Create endpoints
	group.POST("", func(c *gin.Context) {
		controllers.Add(c)
	})
	group.DELETE(":batchExportID", func(c *gin.Context) {
		controllers.Delete(c)
	})
	group.GET("", func(c *gin.Context) {
		controllers.GetAllPaginated(c)
	})
	group.GET(":batchExportID", func(c *gin.Context) {
		controllers.Get(c)
	})
	group.POST("/query", func(c *gin.Context) {
		controllers.Query(c)
	})
	group.PUT(":batchExportID", func(c *gin.Context) {
		controllers.Update(c)
	})


	// Initialize custom routes
	InitCustomRoutes(r)
}
