package product_controllers

import (
	"net/http"
	product_operations "github.com/bakeable/bkry/internal/server/transport/operations/product"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get product ID from URL
	productID := c.Param("productID")

	// Get product
	entity := product_operations.Get(getRepository(), productID)

	// Return product
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
