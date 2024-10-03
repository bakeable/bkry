package product_package

import "encoding/json"

// ToMap converts a ProductPackage struct to a map
func ToMap(e ProductPackage) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a ProductPackage struct to a JSON string
func ToJSON(e ProductPackage) string {
    data, _ := json.Marshal(e)
    return string(data)
}