package foreign

// Convert database map to Key struct
func DecodeKey(m map[string]interface{}) Key {
	return Key{
		ID: m["id"].(string),
		DocumentPath: m["documentPath"].(string),
	}
}

// Convert database map to Key struct
func DecodeKeys(ms []map[string]interface{}) Keys {
	keys := Keys{}
	for _, m := range ms {
		keys = append(keys, DecodeKey(m))
	}
	return keys
}


// Convert database map to ActivationKey struct
func DecodeActivationKey(m map[string]interface{}) ActivationKey {
	return ActivationKey{
		ID: m["id"].(string),
		DocumentPath: m["documentPath"].(string),
		Active: m["active"].(bool),
	}
}

// Convert database map to ActivationKey struct
func DecodeActivationKeys(ms []map[string]interface{}) ActivationKeys {
	keys := ActivationKeys{}
	for _, m := range ms {
		keys = append(keys, DecodeActivationKey(m))
	}
	return keys
}

// Decode a specific key in the foreign key map
func DecodeSingleKey(m map[string]interface{}, key string) Key {
	if _, ok := m[key]; ok {
		if k, ok := m[key]; ok {
			return DecodeKey(k.(map[string]interface{}))
		}
	}
	return Key{}
}

// Decode multiple keys in the foreign key map
func DecodeMultipleKeys(m map[string]interface{}, key string) Keys {
	if _, ok := m[key]; ok {
		if k, ok := m[key].([]interface{}); ok {
			keys := Keys{}
			for _, key := range k {
				keys = append(keys, DecodeKey(key.(map[string]interface{})))
			}
			return keys
		}
	}
	return Keys{}
}

// Decode a specific activation key in the foreign key map
func DecodeSingleActivationKey(m map[string]interface{}, key string) ActivationKey {
	if _, ok := m[key]; ok {
		if k, ok := m[key]; ok {
			return DecodeActivationKey(k.(map[string]interface{}))
		}
	}
	return ActivationKey{}
}

// Decode multiple activation keys in the foreign key map
func DecodeMultipleActivationKeys(m map[string]interface{}, key string) ActivationKeys {
	if _, ok := m[key]; ok {
		if k, ok := m[key].([]interface{}); ok {
			keys := ActivationKeys{}
			for _, key := range k {
				keys = append(keys, DecodeActivationKey(key.(map[string]interface{})))
			}
			return keys
		}
	}
	return ActivationKeys{}
}