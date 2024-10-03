package generator

import "encoding/json"

func notIn(item string, list []string) bool {
    for _, b := range list {
        if b == item {
            return false
        }
    }
    return true
}

func alreadyGenerated(name string, generated map[string]bool) bool {
    if _, ok := generated[name]; ok {
        return true
    }
    generated[name] = true
    return false
}

func hasItems(input interface{}) bool {
    data, err := json.Marshal(input)
    if err != nil {
        return false
    }

    var list []interface{}
    err = json.Unmarshal(data, &list)
    if err != nil {
        return false
    }
    
    return len(list) > 0
}