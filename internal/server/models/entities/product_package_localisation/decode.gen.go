package product_package_localisation

import "encoding/json"

// Decode converts a map to ProductPackageLocalisation struct
func Decode(m interface{}) ProductPackageLocalisation {
	var entity ProductPackageLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to ProductPackageLocalisation struct
func DecodeJSON(s string) ProductPackageLocalisation {
	var entity ProductPackageLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of ProductPackageLocalisation struct
func DecodeAll(ms interface{}) []ProductPackageLocalisation {
	var entities []ProductPackageLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
