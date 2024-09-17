package printing_order_supplier_controllers

import (
	"net/http"
	printing_order_supplier_operations "github.com/bakeable/bkry/data/entities/printing_order_supplier/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAll(c *gin.Context) {

	// Get printing_order_supplier
	entities := printing_order_supplier_operations.GetAll(getRepository())

	// Return printing_order_supplier
	c.JSON(http.StatusOK, gin.H{ "items": entities})

}