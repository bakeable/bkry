package attribute_option_set

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AttributeOptionSet from the database
func Get(id string) (AttributeOptionSet, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AttributeOptionSet", id, nil)
    if err != nil {
        return AttributeOptionSet{}, err
    }

    // Convert the entity to a AttributeOptionSet
    return Decode(entity), nil
}

func (e AttributeOptionSet) Get(id string) (AttributeOptionSet, error) {
    return Get(id)
}

// Update updates a AttributeOptionSet in the database
func Update(e AttributeOptionSet) error {
    // Convert the AttributeOptionSet to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AttributeOptionSet", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AttributeOptionSet) Update() error {
    return Update(e)
}