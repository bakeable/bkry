package portal_event

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a PortalEvent from the database
func Get(id string) (PortalEvent, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("PortalEvent", id, nil)
    if err != nil {
        return PortalEvent{}, err
    }

    // Convert the entity to a PortalEvent
    return Decode(entity), nil
}

func (e PortalEvent) Get(id string) (PortalEvent, error) {
    return Get(id)
}

// Update updates a PortalEvent in the database
func Update(e PortalEvent) error {
    // Convert the PortalEvent to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("PortalEvent", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e PortalEvent) Update() error {
    return Update(e)
}