package airline_pricing

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AirlinePricing from the database
func Get(id string) (AirlinePricing, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AirlinePricing", id, nil)
    if err != nil {
        return AirlinePricing{}, err
    }

    // Convert the entity to a AirlinePricing
    return Decode(entity), nil
}

func (e AirlinePricing) Get(id string) (AirlinePricing, error) {
    return Get(id)
}

// Update updates a AirlinePricing in the database
func Update(e AirlinePricing) error {
    // Convert the AirlinePricing to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AirlinePricing", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AirlinePricing) Update() error {
    return Update(e)
}