package product_controllers

import (
	"net/http"
	product "github.com/bakeable/bkry/internal/server/models/entities/product"
	product_operations "github.com/bakeable/bkry/internal/server/transport/operations/product"
	"github.com/bakeable/bkry/tools"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Add(c *gin.Context) {


	// Get Product mapping from request body
	var mapping map[string]interface{}
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode JSON to Product
	entity := product.Decode(mapping)

	// Create Product
	id := product_operations.Save(getRepository(), entity, utils.GetEditorId(c))

	// Return collected item ID
	c.JSON(http.StatusCreated, gin.H{"id": id})

}