package user

import "encoding/json"

// Decode converts a map to User struct
func Decode(m interface{}) User {
	var entity User
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to User struct
func DecodeJSON(s string) User {
	var entity User
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of User struct
func DecodeAll(ms interface{}) []User {
	var entities []User
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
