package purchase

import "encoding/json"

// Decode converts a map to Purchase struct
func Decode(m interface{}) Purchase {
	var entity Purchase
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to Purchase struct
func DecodeJSON(s string) Purchase {
	var entity Purchase
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of Purchase struct
func DecodeAll(ms interface{}) []Purchase {
	var entities []Purchase
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
