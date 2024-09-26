package airline_settings

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AirlineSettings from the database
func Get(id string) (AirlineSettings, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AirlineSettings", id, nil)
    if err != nil {
        return AirlineSettings{}, err
    }

    // Convert the entity to a AirlineSettings
    return Decode(entity), nil
}

func (e AirlineSettings) Get(id string) (AirlineSettings, error) {
    return Get(id)
}

// Update updates a AirlineSettings in the database
func Update(e AirlineSettings) error {
    // Convert the AirlineSettings to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AirlineSettings", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AirlineSettings) Update() error {
    return Update(e)
}