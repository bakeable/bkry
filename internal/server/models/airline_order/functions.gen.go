package airline_order

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AirlineOrder from the database
func Get(id string) (AirlineOrder, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AirlineOrder", id, nil)
    if err != nil {
        return AirlineOrder{}, err
    }

    // Convert the entity to a AirlineOrder
    return Decode(entity), nil
}

func (e AirlineOrder) Get(id string) (AirlineOrder, error) {
    return Get(id)
}

// Update updates a AirlineOrder in the database
func Update(e AirlineOrder) error {
    // Convert the AirlineOrder to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AirlineOrder", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AirlineOrder) Update() error {
    return Update(e)
}