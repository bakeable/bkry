package airline_pricing_controllers

import (
	"net/http"
	airline_pricing "github.com/bakeable/bkry/internal/server/models/entities/airline_pricing"
	airline_pricing_operations "github.com/bakeable/bkry/internal/server/transport/operations/airline_pricing"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Update(c *gin.Context) {


	// Get AirlinePricing mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to AirlinePricing
	entity := airline_pricing.Decode(mapping)

	// Add AirlinePricing to event
	airline_pricing_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return successfull
	c.Status(http.StatusNoContent)
}