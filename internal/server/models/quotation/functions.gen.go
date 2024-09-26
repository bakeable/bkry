package quotation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Quotation from the database
func Get(id string) (Quotation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Quotation", id, nil)
    if err != nil {
        return Quotation{}, err
    }

    // Convert the entity to a Quotation
    return Decode(entity), nil
}

func (e Quotation) Get(id string) (Quotation, error) {
    return Get(id)
}

// Update updates a Quotation in the database
func Update(e Quotation) error {
    // Convert the Quotation to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Quotation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Quotation) Update() error {
    return Update(e)
}