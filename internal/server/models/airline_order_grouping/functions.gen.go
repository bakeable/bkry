package airline_order_grouping

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AirlineOrderGrouping from the database
func Get(id string) (AirlineOrderGrouping, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AirlineOrderGrouping", id, nil)
    if err != nil {
        return AirlineOrderGrouping{}, err
    }

    // Convert the entity to a AirlineOrderGrouping
    return Decode(entity), nil
}

func (e AirlineOrderGrouping) Get(id string) (AirlineOrderGrouping, error) {
    return Get(id)
}

// Update updates a AirlineOrderGrouping in the database
func Update(e AirlineOrderGrouping) error {
    // Convert the AirlineOrderGrouping to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AirlineOrderGrouping", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AirlineOrderGrouping) Update() error {
    return Update(e)
}