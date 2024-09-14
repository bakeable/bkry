package fail

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DuplicationError struct {
	// ID is the ID of the duplicated resource
	ID string `json:"id"`
	// Path is the path to the duplicated resource
	Path string `json:"path"`
}

// GetResponse returns the response for a DuplicationError
func (e DuplicationError) GetData() map[string]interface{} {
	// Create response
	response := make(map[string]interface{})

	// Set properties
	response["id"] = e.ID
	response["path"] = e.Path

	// Return response
	return response

}

// GetError returns the error type for a DuplicationError
func (e DuplicationError) GetError() string {
	return Duplication
}

// GetResponseCode returns the response code for a DuplicationError
func (e DuplicationError) GetResponseCode() int {
	return http.StatusConflict
}


// ToError returns the DuplicationError as an error
func (e DuplicationError) ToError() error {
	// Convert map[string]interface{} to string
	data, err := json.Marshal(e.GetData())
	if err != nil {
		return err
	}

	return fmt.Errorf("fail:" + e.GetError() + ":" + string(data))
}


// NewDuplicationError creates a new DuplicationError
func NewDuplicationError(id string, path string) IError {
	return DuplicationError{
		ID: id,
		Path: path,
	}
}