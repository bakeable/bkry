package product

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Product from the database
func Get(id string) (Product, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Product", id, nil)
    if err != nil {
        return Product{}, err
    }

    // Convert the entity to a Product
    return Decode(entity), nil
}

func (e Product) Get(id string) (Product, error) {
    return Get(id)
}

// Update updates a Product in the database
func Update(e Product) error {
    // Convert the Product to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Product", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Product) Update() error {
    return Update(e)
}