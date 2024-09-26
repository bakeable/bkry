package printing_order_item

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PrintingOrderItem from the database
func Get(id string) (PrintingOrderItem, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PrintingOrderItem", id, nil)
    if err != nil {
        return PrintingOrderItem{}, err
    }

    // Convert the entity to a PrintingOrderItem
    return Decode(entity), nil
}

func (e PrintingOrderItem) Get(id string) (PrintingOrderItem, error) {
    return Get(id)
}

// Update updates a PrintingOrderItem in the database
func Update(e PrintingOrderItem) error {
    // Convert the PrintingOrderItem to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PrintingOrderItem", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PrintingOrderItem) Update() error {
    return Update(e)
}