package user_profile

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a UserProfile from the database
func Get(id string) (UserProfile, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("UserProfile", id, nil)
    if err != nil {
        return UserProfile{}, err
    }

    // Convert the entity to a UserProfile
    return Decode(entity), nil
}

func (e UserProfile) Get(id string) (UserProfile, error) {
    return Get(id)
}

// Update updates a UserProfile in the database
func Update(e UserProfile) error {
    // Convert the UserProfile to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("UserProfile", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e UserProfile) Update() error {
    return Update(e)
}