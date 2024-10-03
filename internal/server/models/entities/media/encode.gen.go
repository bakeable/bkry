package media

import "encoding/json"

// ToMap converts a Media struct to a map
func ToMap(e Media) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a Media struct to a JSON string
func ToJSON(e Media) string {
    data, _ := json.Marshal(e)
    return string(data)
}