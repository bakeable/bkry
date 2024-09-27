package price_configuration

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PriceConfiguration from the database
func Get(id string) (PriceConfiguration, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PriceConfiguration", id, nil)
    if err != nil {
        return PriceConfiguration{}, err
    }

    // Convert the entity to a PriceConfiguration
    return Decode(entity), nil
}

func (e PriceConfiguration) Get(id string) (PriceConfiguration, error) {
    return Get(id)
}

// Update updates a PriceConfiguration in the database
func Update(e PriceConfiguration) error {
    // Convert the PriceConfiguration to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PriceConfiguration", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PriceConfiguration) Update() error {
    return Update(e)
}