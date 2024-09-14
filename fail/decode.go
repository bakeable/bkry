package fail

import (
	"encoding/json"
	"strings"
)


func FromError(err error) (IError, error) {
	parts := strings.Split(err.Error(), ":")
	
	// Check if error is a ValidationError
	if parts[0] != "fail" {
		return nil, err
	}

	// Get error type
	var e IError
	switch parts[1] {
	case Validation:
		e = ValidationError{}
	case NotFound:
		e = NotFoundError{}
	default:
		return nil, err
	}

	// Convert error to map[string]interface{}
	err = json.Unmarshal([]byte(parts[2]), &e)
	if err != nil {
		return nil, err
	}

	return e, err
}
