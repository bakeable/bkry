package author

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Author from the database
func Get(id string) (Author, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Author", id, nil)
    if err != nil {
        return Author{}, err
    }

    // Convert the entity to a Author
    return Decode(entity), nil
}

func (e Author) Get(id string) (Author, error) {
    return Get(id)
}

// Update updates a Author in the database
func Update(e Author) error {
    // Convert the Author to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Author", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Author) Update() error {
    return Update(e)
}