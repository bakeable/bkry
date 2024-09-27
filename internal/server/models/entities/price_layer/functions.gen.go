package price_layer

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PriceLayer from the database
func Get(id string) (PriceLayer, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PriceLayer", id, nil)
    if err != nil {
        return PriceLayer{}, err
    }

    // Convert the entity to a PriceLayer
    return Decode(entity), nil
}

func (e PriceLayer) Get(id string) (PriceLayer, error) {
    return Get(id)
}

// Update updates a PriceLayer in the database
func Update(e PriceLayer) error {
    // Convert the PriceLayer to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PriceLayer", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PriceLayer) Update() error {
    return Update(e)
}