package airline_settings_controllers

import (
	"net/http"
	airline_settings "github.com/bakeable/bkry/internal/server/models/entities/airline_settings"
	airline_settings_operations "github.com/bakeable/bkry/internal/server/transport/operations/airline_settings"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Update(c *gin.Context) {


	// Get AirlineSettings mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to AirlineSettings
	entity := airline_settings.Decode(mapping)

	// Add AirlineSettings to event
	airline_settings_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return successfull
	c.Status(http.StatusNoContent)
}