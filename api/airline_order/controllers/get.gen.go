package airline_order_controllers

import (
	"net/http"
	airline_order_operations "github.com/bakeable/bkry/data/entities/airline_order/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get airline_order ID from URL
	airlineOrderID := c.Param("airlineOrderID")

	// Get airline_order
	entity := airline_order_operations.Get(getRepository(), airlineOrderID)

	// Return airline_order
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
