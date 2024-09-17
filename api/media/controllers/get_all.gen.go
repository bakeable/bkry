package media_controllers

import (
	"net/http"
	media_operations "github.com/bakeable/bkry/data/entities/media/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get media
	entities := media_operations.GetAll(getRepository())

	// Return media
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}