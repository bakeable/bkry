package examination_settings_controllers

import (
	"net/http"
	examination_settings "github.com/bakeable/bkry/internal/server/models/entities/examination_settings"
	examination_settings_operations "github.com/bakeable/bkry/internal/server/transport/operations/examination_settings"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Add(c *gin.Context) {


	// Get ExaminationSettings mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to ExaminationSettings
	entity := examination_settings.Decode(mapping)

	// Create ExaminationSettings
	id := examination_settings_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return collected item ID
	c.JSON(http.StatusCreated, gin.H{"id": id})

}