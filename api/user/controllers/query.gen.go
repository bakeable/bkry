package user_controllers

import (
	user_operations "github.com/bakeable/bkry/data/entities/user/operations"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	"github.com/bakeable/bkry/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

type QueryBody struct {
	Queries    []datastore.Query    `json:"queries"`
	Pagination datastore.Pagination `json:"pagination"`
}

func Query(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%v", r)
			url := utils.ExtractURLFromErrorMessage(errMsg)
			if url == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
				return
			}
			c.JSON(http.StatusExpectationFailed, gin.H{"error": errMsg, "url": url})
		}
	}()


	// Parse queries and format as []datastore.Query{}
	var queryBody QueryBody
	if err := c.ShouldBindJSON(&queryBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Get users from queries
	entities := user_operations.Query(getRepository(), queryBody.Queries, queryBody.Pagination)

	// Return user
	c.JSON(http.StatusOK, gin.H{"items": entities, "pagination": queryBody.Pagination })
}