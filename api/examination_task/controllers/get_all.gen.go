package examination_task_controllers

import (
	"net/http"
	examination_task_operations "github.com/bakeable/bkry/data/entities/examination_task/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get examination_task
	entities := examination_task_operations.GetAll(getRepository())

	// Return examination_task
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}