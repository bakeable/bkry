package email

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a Email from the database
func Get(id string) (Email, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("Email", id, nil)
    if err != nil {
        return Email{}, err
    }

    // Convert the entity to a Email
    return Decode(entity), nil
}

func (e Email) Get(id string) (Email, error) {
    return Get(id)
}

// Update updates a Email in the database
func Update(e Email) error {
    // Convert the Email to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("Email", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e Email) Update() error {
    return Update(e)
}