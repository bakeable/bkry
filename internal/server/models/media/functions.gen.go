package media

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Media from the database
func Get(id string) (Media, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Media", id, nil)
    if err != nil {
        return Media{}, err
    }

    // Convert the entity to a Media
    return Decode(entity), nil
}

func (e Media) Get(id string) (Media, error) {
    return Get(id)
}

// Update updates a Media in the database
func Update(e Media) error {
    // Convert the Media to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Media", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Media) Update() error {
    return Update(e)
}