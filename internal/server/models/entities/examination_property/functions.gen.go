package examination_property

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ExaminationProperty from the database
func Get(id string) (ExaminationProperty, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ExaminationProperty", id, nil)
    if err != nil {
        return ExaminationProperty{}, err
    }

    // Convert the entity to a ExaminationProperty
    return Decode(entity), nil
}

func (e ExaminationProperty) Get(id string) (ExaminationProperty, error) {
    return Get(id)
}

// Update updates a ExaminationProperty in the database
func Update(e ExaminationProperty) error {
    // Convert the ExaminationProperty to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ExaminationProperty", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ExaminationProperty) Update() error {
    return Update(e)
}