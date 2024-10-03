package question_bundle_localisation

import "github.com/bakeable/bkry/third_party/gcloud/datastore"

// Get retrieves a QuestionBundleLocalisation from the database
func Get(id string) (QuestionBundleLocalisation, error) {
    // Get the data from the database
    entity, err := datastore.GetEntityById("QuestionBundleLocalisation", id, nil)
    if err != nil {
        return QuestionBundleLocalisation{}, err
    }

    // Convert the entity to a QuestionBundleLocalisation
    return Decode(entity), nil
}

func (e QuestionBundleLocalisation) Get(id string) (QuestionBundleLocalisation, error) {
    return Get(id)
}

// Update updates a QuestionBundleLocalisation in the database
func Update(e QuestionBundleLocalisation) error {
    // Convert the QuestionBundleLocalisation to an entity
    entity := ToMap(e)

    // Update the entity in the database
    _, err := datastore.UpdateEntityById("QuestionBundleLocalisation", e.ID, entity, nil)
    if err != nil {
        return err
    }

    return nil
}

func (e QuestionBundleLocalisation) Update() error {
    return Update(e)
}