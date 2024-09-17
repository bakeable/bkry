package quotation_controllers

import (
	"net/http"
	quotation_operations "github.com/bakeable/bkry/data/entities/quotation/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Delete(c *gin.Context) {

	// Get quotation ID from URL
	quotationID := c.Param("quotationID")

	// Delete Quotation
	quotation_operations.Delete(getRepository(), quotationID)

	// Send response
	c.Status(http.StatusNoContent)

}

