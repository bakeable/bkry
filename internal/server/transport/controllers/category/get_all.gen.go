package category_controllers

import (
	"net/http"
	category_operations "github.com/bakeable/bkry/internal/server/transport/operations/category"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get category
	entities := category_operations.GetAll(getRepository())

	// Return category
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}