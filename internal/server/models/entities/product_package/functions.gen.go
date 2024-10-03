package product_package

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ProductPackage from the database
func Get(id string) (ProductPackage, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ProductPackage", id, nil)
    if err != nil {
        return ProductPackage{}, err
    }

    // Convert the entity to a ProductPackage
    return Decode(entity), nil
}

func (e ProductPackage) Get(id string) (ProductPackage, error) {
    return Get(id)
}

// Update updates a ProductPackage in the database
func Update(e ProductPackage) error {
    // Convert the ProductPackage to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ProductPackage", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ProductPackage) Update() error {
    return Update(e)
}