package game_modus

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a GameModus from the database
func Get(id string) (GameModus, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("GameModus", id, nil)
    if err != nil {
        return GameModus{}, err
    }

    // Convert the entity to a GameModus
    return Decode(entity), nil
}

func (e GameModus) Get(id string) (GameModus, error) {
    return Get(id)
}

// Update updates a GameModus in the database
func Update(e GameModus) error {
    // Convert the GameModus to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("GameModus", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e GameModus) Update() error {
    return Update(e)
}