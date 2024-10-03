package question_context_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a QuestionContextLocalisation from the database
func Get(id string) (QuestionContextLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("QuestionContextLocalisation", id, nil)
    if err != nil {
        return QuestionContextLocalisation{}, err
    }

    // Convert the entity to a QuestionContextLocalisation
    return Decode(entity), nil
}

func (e QuestionContextLocalisation) Get(id string) (QuestionContextLocalisation, error) {
    return Get(id)
}

// Update updates a QuestionContextLocalisation in the database
func Update(e QuestionContextLocalisation) error {
    // Convert the QuestionContextLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("QuestionContextLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e QuestionContextLocalisation) Update() error {
    return Update(e)
}