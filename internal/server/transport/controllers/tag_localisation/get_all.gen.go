package tag_localisation_controllers

import (
	"net/http"
	tag_localisation_operations "github.com/bakeable/bkry/internal/server/transport/operations/tag_localisation"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get tag ID from URL
	tagID := c.Param("tagID")

	// Get tag_localisation
	entities := tag_localisation_operations.GetAll(getRepository(), tagID)

	// Return tag_localisation
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}