package game_modus

import "encoding/json"

// ToMap converts a GameModus struct to a map
func ToMap(e GameModus) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a GameModus struct to a JSON string
func ToJSON(e GameModus) string {
    data, _ := json.Marshal(e)
    return string(data)
}