package delivery_entry

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a DeliveryEntry from the database
func Get(id string) (DeliveryEntry, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("DeliveryEntry", id, nil)
    if err != nil {
        return DeliveryEntry{}, err
    }

    // Convert the entity to a DeliveryEntry
    return Decode(entity), nil
}

func (e DeliveryEntry) Get(id string) (DeliveryEntry, error) {
    return Get(id)
}

// Update updates a DeliveryEntry in the database
func Update(e DeliveryEntry) error {
    // Convert the DeliveryEntry to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("DeliveryEntry", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e DeliveryEntry) Update() error {
    return Update(e)
}