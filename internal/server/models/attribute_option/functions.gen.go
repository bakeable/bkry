package attribute_option

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a AttributeOption from the database
func Get(id string) (AttributeOption, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("AttributeOption", id, nil)
    if err != nil {
        return AttributeOption{}, err
    }

    // Convert the entity to a AttributeOption
    return Decode(entity), nil
}

func (e AttributeOption) Get(id string) (AttributeOption, error) {
    return Get(id)
}

// Update updates a AttributeOption in the database
func Update(e AttributeOption) error {
    // Convert the AttributeOption to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("AttributeOption", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e AttributeOption) Update() error {
    return Update(e)
}