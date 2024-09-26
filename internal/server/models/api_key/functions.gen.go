package api_key

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ApiKey from the database
func Get(id string) (ApiKey, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ApiKey", id, nil)
    if err != nil {
        return ApiKey{}, err
    }

    // Convert the entity to a ApiKey
    return Decode(entity), nil
}

func (e ApiKey) Get(id string) (ApiKey, error) {
    return Get(id)
}

// Update updates a ApiKey in the database
func Update(e ApiKey) error {
    // Convert the ApiKey to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ApiKey", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ApiKey) Update() error {
    return Update(e)
}