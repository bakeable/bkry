package attribute_option_controllers

import (
	"net/http"
	attribute_option_operations "github.com/bakeable/bkry/internal/server/transport/operations/attribute_option"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get attribute_option
	entities := attribute_option_operations.GetAll(getRepository())

	// Return attribute_option
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}