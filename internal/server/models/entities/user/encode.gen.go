package user

import "encoding/json"

// ToMap converts a User struct to a map
func ToMap(e User) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a User struct to a JSON string
func ToJSON(e User) string {
    data, _ := json.Marshal(e)
    return string(data)
}