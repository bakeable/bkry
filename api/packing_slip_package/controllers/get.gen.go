package packing_slip_package_controllers

import (
	"net/http"
	packing_slip_package_operations "github.com/bakeable/bkry/data/entities/packing_slip_package/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get packing_slip_package ID from URL
	packingSlipPackageID := c.Param("packingSlipPackageID")

	// Get packing_slip_package
	entity := packing_slip_package_operations.Get(getRepository(), packingSlipPackageID)

	// Return packing_slip_package
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
