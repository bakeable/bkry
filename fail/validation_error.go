package fail

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValidationError struct {
	// Properties is a list of properties that failed validation
	Properties []ValidationErrorProperty
}

type ValidationErrorProperty struct {
	// Path is the path to the property that failed validation within the validated object
	Path string

	// Messages contain the messages that describe the validation error(s)
	Messages []string
}

// GetResponse returns the response for a ValidationError
func (e ValidationError) GetData() map[string]interface{} {
	// Create response
	response := make(map[string]interface{})

	// Set properties
	response["properties"] = e.Properties

	// Return response
	return response

}

// GetError returns the error type for a ValidationError
func (e ValidationError) GetError() string {
	return Validation
}

// GetResponseCode returns the response code for a ValidationError
func (e ValidationError) GetResponseCode() int {
	return http.StatusBadRequest
}


// ToError returns the ValidationError as an error
func (e ValidationError) ToError() error {
	// Convert map[string]interface{} to string
	data, err := json.Marshal(e.GetData())
	if err != nil {
		return err
	}

	return fmt.Errorf("fail:" + e.GetError() + ":" + string(data))
}


// NewValidationError creates a new ValidationError
func NewValidationError(properties []ValidationErrorProperty) IError {
	return ValidationError{
		Properties: properties,
	}
}

func NewValidationErrorProperty(path string, messages []string) ValidationErrorProperty {
	return ValidationErrorProperty{
		Path: path,
		Messages: messages,
	}
}
