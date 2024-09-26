package packing_slip

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PackingSlip from the database
func Get(id string) (PackingSlip, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PackingSlip", id, nil)
    if err != nil {
        return PackingSlip{}, err
    }

    // Convert the entity to a PackingSlip
    return Decode(entity), nil
}

func (e PackingSlip) Get(id string) (PackingSlip, error) {
    return Get(id)
}

// Update updates a PackingSlip in the database
func Update(e PackingSlip) error {
    // Convert the PackingSlip to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PackingSlip", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PackingSlip) Update() error {
    return Update(e)
}