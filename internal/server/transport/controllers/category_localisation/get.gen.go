package category_localisation_controllers

import (
	"net/http"
	category_localisation_operations "github.com/bakeable/bkry/internal/server/transport/operations/category_localisation"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get category ID from URL
	categoryID := c.Param("categoryID")

	// Get category_localisation ID from URL
	categoryLocalisationID := c.Param("categoryLocalisationID")

	// Get category_localisation
	entity := category_localisation_operations.Get(getRepository(), categoryID, categoryLocalisationID)

	// Return category_localisation
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
