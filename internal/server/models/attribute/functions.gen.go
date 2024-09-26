package attribute

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Attribute from the database
func Get(id string) (Attribute, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Attribute", id, nil)
    if err != nil {
        return Attribute{}, err
    }

    // Convert the entity to a Attribute
    return Decode(entity), nil
}

func (e Attribute) Get(id string) (Attribute, error) {
    return Get(id)
}

// Update updates a Attribute in the database
func Update(e Attribute) error {
    // Convert the Attribute to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Attribute", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Attribute) Update() error {
    return Update(e)
}