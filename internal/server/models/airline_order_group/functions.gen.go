package airline_order_group

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AirlineOrderGroup from the database
func Get(id string) (AirlineOrderGroup, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AirlineOrderGroup", id, nil)
    if err != nil {
        return AirlineOrderGroup{}, err
    }

    // Convert the entity to a AirlineOrderGroup
    return Decode(entity), nil
}

func (e AirlineOrderGroup) Get(id string) (AirlineOrderGroup, error) {
    return Get(id)
}

// Update updates a AirlineOrderGroup in the database
func Update(e AirlineOrderGroup) error {
    // Convert the AirlineOrderGroup to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AirlineOrderGroup", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AirlineOrderGroup) Update() error {
    return Update(e)
}