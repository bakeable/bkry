package user_profile_routes

import (
	user_profile_controllers "github.com/bakeable/bkry/api/user_profile/controllers"
	"github.com/bakeable/bkry/api/middleware"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func getControllers() user_profile_controllers.IControllers {
	return user_profile_controllers.NewControllers()
}

func InitRoutes(r *gin.RouterGroup) {
	// Get controllers
	controllers := getControllers()

	// Create a new route group for the UserProfile entity
	group := r.Group("/user_profile")

	// Add authentication middleware for all endpoints
	group.Use(middleware.AuthenticatedRoutes())
	
	// Create endpoints
	group.POST("", func(c *gin.Context) {
		controllers.Add(c)
	})
	group.DELETE(":userProfileID", func(c *gin.Context) {
		controllers.Delete(c)
	})
	group.GET("", func(c *gin.Context) {
		controllers.GetAllPaginated(c)
	})
	group.GET(":userProfileID", func(c *gin.Context) {
		controllers.Get(c)
	})
	group.POST("/query", func(c *gin.Context) {
		controllers.Query(c)
	})
	group.PUT(":userProfileID", func(c *gin.Context) {
		controllers.Update(c)
	})


	// Initialize custom routes
	InitCustomRoutes(r)
}
