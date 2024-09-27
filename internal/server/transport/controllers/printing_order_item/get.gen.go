package printing_order_item_controllers

import (
	"net/http"
	printing_order_item_operations "github.com/bakeable/bkry/internal/server/transport/operations/printing_order_item"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get printing_order_item ID from URL
	printingOrderItemID := c.Param("printingOrderItemID")

	// Get printing_order_item
	entity := printing_order_item_operations.Get(getRepository(), printingOrderItemID)

	// Return printing_order_item
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
