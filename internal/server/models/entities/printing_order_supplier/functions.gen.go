package printing_order_supplier

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PrintingOrderSupplier from the database
func Get(id string) (PrintingOrderSupplier, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PrintingOrderSupplier", id, nil)
    if err != nil {
        return PrintingOrderSupplier{}, err
    }

    // Convert the entity to a PrintingOrderSupplier
    return Decode(entity), nil
}

func (e PrintingOrderSupplier) Get(id string) (PrintingOrderSupplier, error) {
    return Get(id)
}

// Update updates a PrintingOrderSupplier in the database
func Update(e PrintingOrderSupplier) error {
    // Convert the PrintingOrderSupplier to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PrintingOrderSupplier", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PrintingOrderSupplier) Update() error {
    return Update(e)
}