package user

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a User from the database
func Get(id string) (User, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("User", id, nil)
    if err != nil {
        return User{}, err
    }

    // Convert the entity to a User
    return Decode(entity), nil
}

func (e User) Get(id string) (User, error) {
    return Get(id)
}

// Update updates a User in the database
func Update(e User) error {
    // Convert the User to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("User", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e User) Update() error {
    return Update(e)
}