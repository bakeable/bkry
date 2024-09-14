package fail

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	// Path is the path to the entity that could not be found
	Path string

	// Queries is a list of queries that were used to try and find the entity
	Queries []string
}

// GetResponse returns the response for a NotFoundError
func (e NotFoundError) GetData() map[string]interface{} {
	// Create response
	response := make(map[string]interface{})

	// Set properties
	response["path"] = e.Path
	response["queries"] = e.Queries

	// Return response
	return response

}

// GetError returns the error type for a NotFoundError
func (e NotFoundError) GetError() string {
	return NotFound
}

// GetResponseCode returns the response code for a NotFoundError
func (e NotFoundError) GetResponseCode() int {
	return http.StatusNotFound
}

// ToError returns the NotFoundError as an error
func (e NotFoundError) ToError() error {
	// Convert map[string]interface{} to string
	data, err := json.Marshal(e.GetData())
	if err != nil {
		return err
	}

	return fmt.Errorf("fail:" + e.GetError() + ":" + string(data))
}

// NewNotFoundError creates a new NotFoundError
func NewNotFoundError(path string, queries []string) IError {
	return NotFoundError{
		Path: path,
		Queries: queries,
	}
}