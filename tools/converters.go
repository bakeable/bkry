package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"
)

func ConvertToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return float64(0)
	}
}

// Format ID to binary
func StringToBinary(str string) []byte {
	binaryId, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        fmt.Println("Error decoding Base64:", err)
        panic(err)
    }

	return binaryId
}

// PascalCaseToCamelCase converts camelCase string to PascalCase.
func CamelCaseToPascalCase(input string) string {
	var buffer bytes.Buffer
	for i, r := range input {
		if i == 0 {
			buffer.WriteRune(unicode.ToUpper(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// PascalCaseToCamelCase converts PascalCase string to camelCase.
func PascalCaseToCamelCase(input string) string {
	var buffer bytes.Buffer
	for i, r := range input {
		if i == 0 {
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// PascalCaseToSnakeCase converts PascalCase string to snake_case.
func PascalCaseToSnakeCase(input string) string {
	var buffer bytes.Buffer
	for i, r := range input {
		if i > 0 && unicode.IsUpper(r) {
			buffer.WriteRune('_')
		}
		buffer.WriteRune(unicode.ToLower(r))
	}
	return buffer.String()
}


// SnakeCaseToPascalCase converts snake_case string to PascalCase.
func SnakeCaseToPascalCase(input string) string {
	var buffer bytes.Buffer
	words := strings.Split(input, "_")
	for _, word := range words {
		for i, r := range word {
			if i == 0 {
				buffer.WriteRune(unicode.ToUpper(r))
			} else {
				buffer.WriteRune(r)
			}
		}
	}
	return buffer.String()
}