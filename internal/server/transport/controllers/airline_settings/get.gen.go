package airline_settings_controllers

import (
	"net/http"
	airline_settings_operations "github.com/bakeable/bkry/internal/server/transport/operations/airline_settings"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get airline_settings ID from URL
	airlineSettingsID := c.Param("airlineSettingsID")

	// Get airline_settings
	entity := airline_settings_operations.Get(getRepository(), airlineSettingsID)

	// Return airline_settings
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
