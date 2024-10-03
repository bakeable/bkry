package product_package_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ProductPackageLocalisation from the database
func Get(id string) (ProductPackageLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ProductPackageLocalisation", id, nil)
    if err != nil {
        return ProductPackageLocalisation{}, err
    }

    // Convert the entity to a ProductPackageLocalisation
    return Decode(entity), nil
}

func (e ProductPackageLocalisation) Get(id string) (ProductPackageLocalisation, error) {
    return Get(id)
}

// Update updates a ProductPackageLocalisation in the database
func Update(e ProductPackageLocalisation) error {
    // Convert the ProductPackageLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ProductPackageLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ProductPackageLocalisation) Update() error {
    return Update(e)
}