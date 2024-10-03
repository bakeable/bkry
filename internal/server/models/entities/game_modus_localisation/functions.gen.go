package game_modus_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a GameModusLocalisation from the database
func Get(id string) (GameModusLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("GameModusLocalisation", id, nil)
    if err != nil {
        return GameModusLocalisation{}, err
    }

    // Convert the entity to a GameModusLocalisation
    return Decode(entity), nil
}

func (e GameModusLocalisation) Get(id string) (GameModusLocalisation, error) {
    return Get(id)
}

// Update updates a GameModusLocalisation in the database
func Update(e GameModusLocalisation) error {
    // Convert the GameModusLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("GameModusLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e GameModusLocalisation) Update() error {
    return Update(e)
}