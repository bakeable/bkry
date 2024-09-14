package foreign

// Convert Key struct to database map
func (fk Key) Encode() map[string]interface{} {
	return map[string]interface{}{
		"id": fk.ID,
		"documentPath": fk.DocumentPath,
	}
}

// Convert database map to Key struct
func (keys Keys) Encode() []map[string]interface{} {
	maps := []map[string]interface{}{}
	for _, key := range keys {
		maps = append(maps, key.Encode())
	}
	return maps
}


// Convert Key struct to database map
func (fk ActivationKey) Encode() map[string]interface{} {
	return map[string]interface{}{
		"id": fk.ID,
		"documentPath": fk.DocumentPath,
		"active": fk.Active,
	}
}

// Convert database map to Key struct
func (keys ActivationKeys) Encode() []map[string]interface{} {
	maps := []map[string]interface{}{}
	for _, key := range keys {
		maps = append(maps, key.Encode())
	}
	return maps
}