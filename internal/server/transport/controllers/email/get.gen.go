package email_controllers

import (
	"net/http"
	email_operations "github.com/bakeable/bkry/internal/server/transport/operations/email"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Get(c *gin.Context) {

	// Get email ID from URL
	emailID := c.Param("emailID")

	// Get email
	entity := email_operations.Get(getRepository(), emailID)

	// Return email
	c.JSON(http.StatusOK, gin.H{ "item": entity })
}
