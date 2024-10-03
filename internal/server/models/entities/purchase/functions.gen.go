package purchase

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Purchase from the database
func Get(id string) (Purchase, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Purchase", id, nil)
    if err != nil {
        return Purchase{}, err
    }

    // Convert the entity to a Purchase
    return Decode(entity), nil
}

func (e Purchase) Get(id string) (Purchase, error) {
    return Get(id)
}

// Update updates a Purchase in the database
func Update(e Purchase) error {
    // Convert the Purchase to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Purchase", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Purchase) Update() error {
    return Update(e)
}