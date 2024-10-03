package purchase

import "encoding/json"

// ToMap converts a Purchase struct to a map
func ToMap(e Purchase) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a Purchase struct to a JSON string
func ToJSON(e Purchase) string {
    data, _ := json.Marshal(e)
    return string(data)
}