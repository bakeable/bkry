package generator

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