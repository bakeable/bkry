package category_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a CategoryLocalisation from the database
func Get(id string) (CategoryLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("CategoryLocalisation", id, nil)
    if err != nil {
        return CategoryLocalisation{}, err
    }

    // Convert the entity to a CategoryLocalisation
    return Decode(entity), nil
}

func (e CategoryLocalisation) Get(id string) (CategoryLocalisation, error) {
    return Get(id)
}

// Update updates a CategoryLocalisation in the database
func Update(e CategoryLocalisation) error {
    // Convert the CategoryLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("CategoryLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e CategoryLocalisation) Update() error {
    return Update(e)
}