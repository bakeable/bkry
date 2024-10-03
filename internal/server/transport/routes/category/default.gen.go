package category_routes

import (
	category_controllers "github.com/bakeable/bkry/internal/server/transport/controllers/category"
	"github.com/bakeable/bkry/internal/server/transport/middleware"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func getControllers() category_controllers.IControllers {
	return category_controllers.NewControllers()
}

func InitRoutes(r *gin.RouterGroup) {
	// Get controllers
	controllers := getControllers()

	// Create a new route group for the Category entity
	group := r.Group("/category")

	// Add authentication middleware for all endpoints
	group.Use(middleware.AuthenticatedRoutes())
	
	// Create endpoints
	group.POST("", func(c *gin.Context) {
		controllers.Add(c)
	})
	group.DELETE(":categoryID", func(c *gin.Context) {
		controllers.Delete(c)
	})
	group.GET("", func(c *gin.Context) {
		controllers.GetAllPaginated(c)
	})
	group.GET(":categoryID", func(c *gin.Context) {
		controllers.Get(c)
	})
	group.POST("/query", func(c *gin.Context) {
		controllers.Query(c)
	})
	group.PUT(":categoryID", func(c *gin.Context) {
		controllers.Update(c)
	})


	// Initialize custom routes
	InitCustomRoutes(r)
}
