package examination_task

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a ExaminationTask from the database
func Get(id string) (ExaminationTask, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("ExaminationTask", id, nil)
    if err != nil {
        return ExaminationTask{}, err
    }

    // Convert the entity to a ExaminationTask
    return Decode(entity), nil
}

func (e ExaminationTask) Get(id string) (ExaminationTask, error) {
    return Get(id)
}

// Update updates a ExaminationTask in the database
func Update(e ExaminationTask) error {
    // Convert the ExaminationTask to an entity
    entity := Encode(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("ExaminationTask", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e ExaminationTask) Update() error {
    return Update(e)
}