package game_modus_localisation

import "encoding/json"

// ToMap converts a GameModusLocalisation struct to a map
func ToMap(e GameModusLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a GameModusLocalisation struct to a JSON string
func ToJSON(e GameModusLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}