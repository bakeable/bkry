package user_controllers

import (
	"net/http"
	user_operations "github.com/bakeable/bkry/data/entities/user/operations"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func Delete(c *gin.Context) {

	// Get user ID from URL
	userID := c.Param("userID")

	// Delete User
	user_operations.Delete(getRepository(), userID)

	// Send response
	c.Status(http.StatusNoContent)

}

