package cloud_function_controllers

import (
	"net/http"
	cloud_function "github.com/bakeable/bkry/data/entities/cloud_function"
	cloud_function_operations "github.com/bakeable/bkry/data/entities/cloud_function/operations"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Update(c *gin.Context) {


	// Get CloudFunction mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to CloudFunction
	entity := cloud_function.Decode(mapping)

	// Add CloudFunction to event
	cloud_function_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return successfull
	c.Status(http.StatusNoContent)
}