package printing_order_supplier_routes

import (
	printing_order_supplier_controllers "github.com/bakeable/bkry/internal/server/transport/controllers/printing_order_supplier"
	"github.com/bakeable/bkry/internal/server/transport/middleware"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT

func getControllers() printing_order_supplier_controllers.IControllers {
	return printing_order_supplier_controllers.NewControllers()
}

func InitRoutes(r *gin.RouterGroup) {
	// Get controllers
	controllers := getControllers()

	// Create a new route group for the PrintingOrderSupplier entity
	group := r.Group("/printing_order_supplier")

	// Add authentication middleware for all endpoints
	group.Use(middleware.AuthenticatedRoutes())
	
	// Create endpoints
	group.POST("", func(c *gin.Context) {
		controllers.Add(c)
	})
	group.DELETE(":printingOrderSupplierID", func(c *gin.Context) {
		controllers.Delete(c)
	})
	group.GET("", func(c *gin.Context) {
		controllers.GetAllPaginated(c)
	})
	group.GET(":printingOrderSupplierID", func(c *gin.Context) {
		controllers.Get(c)
	})
	group.POST("/query", func(c *gin.Context) {
		controllers.Query(c)
	})
	group.PUT(":printingOrderSupplierID", func(c *gin.Context) {
		controllers.Update(c)
	})


	// Initialize custom routes
	InitCustomRoutes(r)
}
