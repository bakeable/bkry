package packing_slip_package

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PackingSlipPackage from the database
func Get(id string) (PackingSlipPackage, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PackingSlipPackage", id, nil)
    if err != nil {
        return PackingSlipPackage{}, err
    }

    // Convert the entity to a PackingSlipPackage
    return Decode(entity), nil
}

func (e PackingSlipPackage) Get(id string) (PackingSlipPackage, error) {
    return Get(id)
}

// Update updates a PackingSlipPackage in the database
func Update(e PackingSlipPackage) error {
    // Convert the PackingSlipPackage to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PackingSlipPackage", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PackingSlipPackage) Update() error {
    return Update(e)
}