package tag

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Tag from the database
func Get(id string) (Tag, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Tag", id, nil)
    if err != nil {
        return Tag{}, err
    }

    // Convert the entity to a Tag
    return Decode(entity), nil
}

func (e Tag) Get(id string) (Tag, error) {
    return Get(id)
}

// Update updates a Tag in the database
func Update(e Tag) error {
    // Convert the Tag to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Tag", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Tag) Update() error {
    return Update(e)
}