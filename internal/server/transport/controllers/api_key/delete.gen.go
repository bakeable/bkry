package api_key_controllers

import (
	"net/http"
	api_key_operations "github.com/bakeable/bkry/internal/server/transport/operations/api_key"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Delete(c *gin.Context) {

	// Get api_key ID from URL
	apiKeyID := c.Param("apiKeyID")

	// Delete ApiKey
	api_key_operations.Delete(getRepository(), apiKeyID)

	// Send response
	c.Status(http.StatusNoContent)

}

