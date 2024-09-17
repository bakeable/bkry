package price_configuration_controllers

import (
	"net/http"
	price_configuration_operations "github.com/bakeable/bkry/data/entities/price_configuration/operations"
	"strconv"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	

	"github.com/gin-gonic/gin"
)

//// THIS FILE IS AUTOGENERATED. DO NOT EDIT
//// IF YOU WISH TO EDIT THIS CONTROLLER, REMOVE THE .gen.go EXTENSION
//// THIS WAY, IT WON'T BE OVERWRITTEN AND THERE WON'T BE A CONTROLLER GENERATED

func GetAllPaginated(c *gin.Context) {


	// Fallback on GetAll if no pagination parameters are provided
	if c.DefaultQuery("page_size", "") == "" || c.DefaultQuery("page_number", "") == "" {
		GetAll(c)
		return
	}

	// Get page size
	pageSizeStr := c.DefaultQuery("page_size", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 1
	}

	// Get page number
	pageNumberStr := c.DefaultQuery("page_number", "1")
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		pageNumber = 1
	}

	// Get ordering parameters
	orderBy := c.DefaultQuery("order_by", "id")
	ascending := c.DefaultQuery("asc", "true") == "true"

	// Get price_configuration
	entities, pagination := price_configuration_operations.GetAllPaginated(getRepository(), datastore.Pagination{
		PageSize:    pageSize,
		PageNumber:  pageNumber,
		OrderBy:     orderBy,
		Ascending:   ascending,
	})

	// Return price_configuration
	c.JSON(http.StatusOK, gin.H{ "items": entities, "pagination": pagination })

}