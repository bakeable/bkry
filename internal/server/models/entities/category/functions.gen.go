package category

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Category from the database
func Get(id string) (Category, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Category", id, nil)
    if err != nil {
        return Category{}, err
    }

    // Convert the entity to a Category
    return Decode(entity), nil
}

func (e Category) Get(id string) (Category, error) {
    return Get(id)
}

// Update updates a Category in the database
func Update(e Category) error {
    // Convert the Category to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Category", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Category) Update() error {
    return Update(e)
}