package product_package

import "encoding/json"

// Decode converts a map to ProductPackage struct
func Decode(m interface{}) ProductPackage {
	var entity ProductPackage
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to ProductPackage struct
func DecodeJSON(s string) ProductPackage {
	var entity ProductPackage
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of ProductPackage struct
func DecodeAll(ms interface{}) []ProductPackage {
	var entities []ProductPackage
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
