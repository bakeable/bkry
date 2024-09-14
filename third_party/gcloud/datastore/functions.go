package datastore

import (
	"fmt"
	"strings"

	"cloud.google.com/go/datastore"
)


func ConvertMapToProperties(data map[string]interface{}) []datastore.Property {
	properties := make([]datastore.Property, 0)
	for key, value := range data {
		properties = append(properties, datastore.Property{Name: key, Value: value})
	}
	return properties
}

func ConvertMapToPropertyList(data map[string]interface{}) datastore.PropertyList {
	properties := make([]datastore.Property, 0)
	for key, value := range data {
		properties = append(properties, datastore.Property{Name: key, Value: value})
	}
	return properties
}


// ConvertPropertiesToMap converts a PropertyList to a map[string]interface{}.
func ConvertPropertiesToMap(properties datastore.PropertyList) map[string]interface{} {
	result := make(map[string]interface{})
	for _, prop := range properties {
		result[prop.Name] = convertValue(prop.Value)
	}
	return result
}

// convertValue handles the conversion of each property value, including nested structures.
func convertValue(value interface{}) interface{} {
	// Check the type of value and handle nested Entity types
	switch v := value.(type) {
	case *datastore.Entity:
		// If the value is a PropertyList (which can be considered as an entity), recursively convert it
		return ConvertPropertiesToMap(v.Properties)
	case []interface{}:
		// If the value is a slice, recursively handle each element
		slice := make([]interface{}, len(v))
		for i, elem := range v {
			slice[i] = convertValue(elem)
		}
		return slice
	case []datastore.Property:
		// If the value is a PropertyList, recursively convert it
		return ConvertPropertiesToMap(v)
	case map[string]interface{}:
		// If the value is a map, recursively convert it
		result := make(map[string]interface{})
		for key, val := range v {
			result[key] = convertValue(val)
		}
		return result
	case map[string]string:
		// If the value is a map, recursively convert it
		result := make(map[string]interface{})
		for key, val := range v {
			result[key] = convertValue(val)
		}
		return result
	case map[int]int:
		// If the value is a map, recursively convert it
		result := make(map[int]interface{})
		for key, val := range v {
			result[key] = convertValue(val)
		}
		return result
	default:
		// For other types (e.g., string, int), return as is
		return v
	}
}

// DeriveKindAndAncestorKey converts a string path to a Datastore key and its ancestor.
func DeriveKindAndAncestorKey(path string) (string, *datastore.Key, error) {
	// Split the path into segments
	segments := strings.Split(path, "/")

	if len(segments) == 1 {
		return segments[0], nil, nil
	}

	if len(segments) == 0 || len(segments)%2 != 0 {
		return "", nil, fmt.Errorf("invalid path format: %s", path)
	}

	// Initialize ancestor key
	var ancestorKey *datastore.Key

	// Iterate through the segments to build the key hierarchy
	for i := 0; i < len(segments)-2; i += 2 {
		kind := segments[i]
		id := segments[i+1]

		if ancestorKey == nil {
			ancestorKey = datastore.NameKey(kind, id, nil)
		} else {
			ancestorKey = datastore.NameKey(kind, id, ancestorKey)
		}
	}

	// The last two segments correspond to the final key
	finalKind := segments[len(segments)-2]

	return finalKind, ancestorKey, nil
}

// DeriveKeyAndAncestorKey converts a string path to a Datastore key and its ancestor.
func DeriveKeyAndAncestorKey(path string) (string, string, *datastore.Key, error) {
	// Split the path into segments
	segments := strings.Split(path, "/")
	
	if len(segments) == 0 || len(segments)%2 != 0 {
		return "", "", nil, fmt.Errorf("invalid path format: %s", path)
	}

	// Initialize ancestor key
	var ancestorKey *datastore.Key

	// Iterate through the segments to build the key hierarchy
	for i := 0; i < len(segments)-2; i += 2 {
		kind := segments[i]
		id := segments[i+1]

		if ancestorKey == nil {
			ancestorKey = datastore.NameKey(kind, id, nil)
		} else {
			ancestorKey = datastore.NameKey(kind, id, ancestorKey)
		}
	}

	// The last two segments correspond to the final key
	finalKind := segments[len(segments)-2]
	finalID := segments[len(segments)-1]

	return finalKind, finalID, ancestorKey, nil
}
