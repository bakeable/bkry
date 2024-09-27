package airline_order_group_controllers

import (
	"net/http"
	airline_order_group "github.com/bakeable/bkry/internal/server/models/entities/airline_order_group"
	airline_order_group_operations "github.com/bakeable/bkry/internal/server/transport/operations/airline_order_group"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Add(c *gin.Context) {


	// Get AirlineOrderGroup mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to AirlineOrderGroup
	entity := airline_order_group.Decode(mapping)

	// Create AirlineOrderGroup
	id := airline_order_group_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return collected item ID
	c.JSON(http.StatusCreated, gin.H{"id": id})

}