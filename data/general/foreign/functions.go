package foreign

// Add a key to an array of keys
func (keys *Keys) Add(key Key) {
	// Check if the foreign key exists
	for _, k := range *keys {
		if k.ID == key.ID {
			return // Key already exists, no need to add it again
		}
	}

	// Add key
	*keys = append(*keys, key)
}


// Remove a key from an array of keys
func (keys *Keys) Remove(id string) {
	for i, k := range *keys {
		if k.ID == id {
			// Remove the foreign key from the array
			*keys = append((*keys)[:i], (*keys)[i+1:]...)
			return
		}
	}
}

// Add an activation key to an array of activation keys
func (keys *ActivationKeys) Add(key ActivationKey) {
	// Check if the foreign key exists
	for _, k := range *keys {
		if k.ID == key.ID {
			return // Key already exists, no need to add it again
		}
	}

	// Add key
	*keys = append(*keys, key)
}
// Remove an activation key from an array of activation keys
func (keys *ActivationKeys) Remove(id string) {
	for i, k := range *keys {
		if k.ID == id {
			// Remove the foreign key from the array
			*keys = append((*keys)[:i], (*keys)[i+1:]...)
			return
		}
	}
}


// Activate an activation key in an array of keys
func (keys ActivationKeys) Activate(id string) ActivationKeys {
	for i, k := range keys {
		if k.ID == id {
			// Activate the foreign key
			keys[i].Active = true
		} else {
			// Deactivate the foreign key
			keys[i].Active = false
		}
	}
	return keys
}