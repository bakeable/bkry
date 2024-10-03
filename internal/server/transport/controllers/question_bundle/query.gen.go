package question_bundle_controllers

import (
	question_bundle_operations "github.com/bakeable/bkry/internal/server/transport/operations/question_bundle"
	"github.com/bakeable/bkry/internal/server/database"
	"github.com/bakeable/bkry/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

type QueryBody struct {
	Queries    []database.Query    `json:"queries"`
	Pagination database.Pagination `json:"pagination"`
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


	// Parse queries and format as []database.Query{}
	var queryBody QueryBody
	if err := c.ShouldBindJSON(&queryBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Get question_bundles from queries
	entities := question_bundle_operations.Query(getRepository(), queryBody.Queries, queryBody.Pagination)

	// Return question_bundle
	c.JSON(http.StatusOK, gin.H{"items": entities, "pagination": queryBody.Pagination })
}