package tag_localisation

import "encoding/json"

// ToMap converts a TagLocalisation struct to a map
func ToMap(e TagLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a TagLocalisation struct to a JSON string
func ToJSON(e TagLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}