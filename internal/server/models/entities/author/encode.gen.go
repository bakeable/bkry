package author

import "encoding/json"

// ToMap converts a Author struct to a map
func ToMap(e Author) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a Author struct to a JSON string
func ToJSON(e Author) string {
    data, _ := json.Marshal(e)
    return string(data)
}