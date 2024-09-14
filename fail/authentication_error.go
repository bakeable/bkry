package fail

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationError struct {}

// GetResponse returns the response for a AuthenticationError
func (e AuthenticationError) GetData() map[string]interface{} {
	// Create response
	response := make(map[string]interface{})

	// Return response
	return response

}

// GetError returns the error type for a AuthenticationError
func (e AuthenticationError) GetError() string {
	return Authentication
}

// GetResponseCode returns the response code for a AuthenticationError
func (e AuthenticationError) GetResponseCode() int {
	return http.StatusBadRequest
}


// ToError returns the AuthenticationError as an error
func (e AuthenticationError) ToError() error {
	// Convert map[string]interface{} to string
	data, err := json.Marshal(e.GetData())
	if err != nil {
		return err
	}

	return fmt.Errorf("fail:" + e.GetError() + ":" + string(data))
}


// NewAuthenticationError creates a new AuthenticationError
func NewAuthenticationError() IError {
	return AuthenticationError{}
}