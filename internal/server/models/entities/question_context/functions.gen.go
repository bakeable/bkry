package question_context

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a QuestionContext from the database
func Get(id string) (QuestionContext, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("QuestionContext", id, nil)
    if err != nil {
        return QuestionContext{}, err
    }

    // Convert the entity to a QuestionContext
    return Decode(entity), nil
}

func (e QuestionContext) Get(id string) (QuestionContext, error) {
    return Get(id)
}

// Update updates a QuestionContext in the database
func Update(e QuestionContext) error {
    // Convert the QuestionContext to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("QuestionContext", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e QuestionContext) Update() error {
    return Update(e)
}