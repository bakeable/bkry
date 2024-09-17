package airline_pricing_controllers

import (
	"net/http"
	airline_pricing_operations "github.com/bakeable/bkry/data/entities/airline_pricing/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get airline_pricing ID from URL
	airlinePricingID := c.Param("airlinePricingID")

	// Get airline_pricing
	entity := airline_pricing_operations.Get(getRepository(), airlinePricingID)

	// Return airline_pricing
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
