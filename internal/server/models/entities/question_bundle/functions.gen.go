package question_bundle

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a QuestionBundle from the database
func Get(id string) (QuestionBundle, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("QuestionBundle", id, nil)
    if err != nil {
        return QuestionBundle{}, err
    }

    // Convert the entity to a QuestionBundle
    return Decode(entity), nil
}

func (e QuestionBundle) Get(id string) (QuestionBundle, error) {
    return Get(id)
}

// Update updates a QuestionBundle in the database
func Update(e QuestionBundle) error {
    // Convert the QuestionBundle to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("QuestionBundle", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e QuestionBundle) Update() error {
    return Update(e)
}