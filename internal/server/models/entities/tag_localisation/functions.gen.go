package tag_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a TagLocalisation from the database
func Get(id string) (TagLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("TagLocalisation", id, nil)
    if err != nil {
        return TagLocalisation{}, err
    }

    // Convert the entity to a TagLocalisation
    return Decode(entity), nil
}

func (e TagLocalisation) Get(id string) (TagLocalisation, error) {
    return Get(id)
}

// Update updates a TagLocalisation in the database
func Update(e TagLocalisation) error {
    // Convert the TagLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("TagLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e TagLocalisation) Update() error {
    return Update(e)
}