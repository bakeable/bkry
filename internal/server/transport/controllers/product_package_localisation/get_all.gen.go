package product_package_localisation_controllers

import (
	"net/http"
	product_package_localisation_operations "github.com/bakeable/bkry/internal/server/transport/operations/product_package_localisation"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get productPackage ID from URL
	productPackageID := c.Param("productPackageID")

	// Get product_package_localisation
	entities := product_package_localisation_operations.GetAll(getRepository(), productPackageID)

	// Return product_package_localisation
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}