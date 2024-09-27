package examination_task_routes

import (
	examination_task_controllers "github.com/bakeable/bkry/internal/server/transport/controllers/examination_task"
	"github.com/bakeable/bkry/internal/server/transport/middleware"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func getControllers() examination_task_controllers.IControllers {
	return examination_task_controllers.NewControllers()
}

func InitRoutes(r *gin.RouterGroup) {
	// Get controllers
	controllers := getControllers()

	// Create a new route group for the ExaminationTask entity
	group := r.Group("/examination_task")

	// Add authentication middleware for all endpoints
	group.Use(middleware.AuthenticatedRoutes())
	
	// Create endpoints
	group.POST("", func(c *gin.Context) {
		controllers.Add(c)
	})
	group.DELETE(":examinationTaskID", func(c *gin.Context) {
		controllers.Delete(c)
	})
	group.GET("", func(c *gin.Context) {
		controllers.GetAllPaginated(c)
	})
	group.GET(":examinationTaskID", func(c *gin.Context) {
		controllers.Get(c)
	})
	group.POST("/query", func(c *gin.Context) {
		controllers.Query(c)
	})
	group.PUT(":examinationTaskID", func(c *gin.Context) {
		controllers.Update(c)
	})


	// Initialize custom routes
	InitCustomRoutes(r)
}
