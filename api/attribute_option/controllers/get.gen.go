package attribute_option_controllers

import (
	"net/http"
	attribute_option_operations "github.com/bakeable/bkry/data/entities/attribute_option/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get attribute_option ID from URL
	attributeOptionID := c.Param("attributeOptionID")

	// Get attribute_option
	entity := attribute_option_operations.Get(getRepository(), attributeOptionID)

	// Return attribute_option
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
