package examination_settings

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ExaminationSettings from the database
func Get(id string) (ExaminationSettings, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ExaminationSettings", id, nil)
    if err != nil {
        return ExaminationSettings{}, err
    }

    // Convert the entity to a ExaminationSettings
    return Decode(entity), nil
}

func (e ExaminationSettings) Get(id string) (ExaminationSettings, error) {
    return Get(id)
}

// Update updates a ExaminationSettings in the database
func Update(e ExaminationSettings) error {
    // Convert the ExaminationSettings to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ExaminationSettings", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ExaminationSettings) Update() error {
    return Update(e)
}