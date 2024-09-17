package airline_order_grouping_controllers

import (
	"net/http"
	airline_order_grouping "github.com/bakeable/bkry/data/entities/airline_order_grouping"
	airline_order_grouping_operations "github.com/bakeable/bkry/data/entities/airline_order_grouping/operations"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Add(c *gin.Context) {


	// Get AirlineOrderGrouping mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to AirlineOrderGrouping
	entity := airline_order_grouping.Decode(mapping)

	// Create AirlineOrderGrouping
	id := airline_order_grouping_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return collected item ID
	c.JSON(http.StatusCreated, gin.H{"id": id})

}