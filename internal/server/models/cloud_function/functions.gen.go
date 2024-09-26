package cloud_function

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a CloudFunction from the database
func Get(id string) (CloudFunction, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("CloudFunction", id, nil)
    if err != nil {
        return CloudFunction{}, err
    }

    // Convert the entity to a CloudFunction
    return Decode(entity), nil
}

func (e CloudFunction) Get(id string) (CloudFunction, error) {
    return Get(id)
}

// Update updates a CloudFunction in the database
func Update(e CloudFunction) error {
    // Convert the CloudFunction to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("CloudFunction", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e CloudFunction) Update() error {
    return Update(e)
}