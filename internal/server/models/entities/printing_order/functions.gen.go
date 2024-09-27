package printing_order

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PrintingOrder from the database
func Get(id string) (PrintingOrder, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PrintingOrder", id, nil)
    if err != nil {
        return PrintingOrder{}, err
    }

    // Convert the entity to a PrintingOrder
    return Decode(entity), nil
}

func (e PrintingOrder) Get(id string) (PrintingOrder, error) {
    return Get(id)
}

// Update updates a PrintingOrder in the database
func Update(e PrintingOrder) error {
    // Convert the PrintingOrder to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PrintingOrder", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PrintingOrder) Update() error {
    return Update(e)
}