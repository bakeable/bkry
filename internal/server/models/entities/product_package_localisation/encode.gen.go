package product_package_localisation

import "encoding/json"

// ToMap converts a ProductPackageLocalisation struct to a map
func ToMap(e ProductPackageLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a ProductPackageLocalisation struct to a JSON string
func ToJSON(e ProductPackageLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}