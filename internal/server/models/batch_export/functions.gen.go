package batch_export

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a BatchExport from the database
func Get(id string) (BatchExport, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("BatchExport", id, nil)
    if err != nil {
        return BatchExport{}, err
    }

    // Convert the entity to a BatchExport
    return Decode(entity), nil
}

func (e BatchExport) Get(id string) (BatchExport, error) {
    return Get(id)
}

// Update updates a BatchExport in the database
func Update(e BatchExport) error {
    // Convert the BatchExport to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("BatchExport", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e BatchExport) Update() error {
    return Update(e)
}