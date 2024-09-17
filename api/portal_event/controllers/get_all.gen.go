package portal_event_controllers

import (
	"net/http"
	portal_event_operations "github.com/bakeable/bkry/data/entities/portal_event/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get portal_event
	entities := portal_event_operations.GetAll(getRepository())

	// Return portal_event
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}