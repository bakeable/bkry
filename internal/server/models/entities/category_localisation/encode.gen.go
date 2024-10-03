package category_localisation

import "encoding/json"

// ToMap converts a CategoryLocalisation struct to a map
func ToMap(e CategoryLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a CategoryLocalisation struct to a JSON string
func ToJSON(e CategoryLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}