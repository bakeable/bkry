package price_configuration_controllers

import (
	"net/http"
	price_configuration_operations "github.com/bakeable/bkry/data/entities/price_configuration/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get price_configuration
	entities := price_configuration_operations.GetAll(getRepository())

	// Return price_configuration
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}