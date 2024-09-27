package examination_property_controllers

import (
	"net/http"
	examination_property_operations "github.com/bakeable/bkry/internal/server/transport/operations/examination_property"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Delete(c *gin.Context) {

	// Get examination_property ID from URL
	examinationPropertyID := c.Param("examinationPropertyID")

	// Delete ExaminationProperty
	examination_property_operations.Delete(getRepository(), examinationPropertyID)

	// Send response
	c.Status(http.StatusNoContent)

}

