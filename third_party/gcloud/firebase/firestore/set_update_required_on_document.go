package db

import utils "github.com/bakeable/bkry/tools"

func SetUpdateRequiredOnDocument(path string, required bool) {
	// Set update required on waste stream
	err := UpdateDocument(path, updateRequiredMap(required), utils.GetEditorId(nil))
	if err != nil {
		panic(handleError(err, "Failed to set update required on document", &path, nil))
	}
}


func updateRequiredMap(required bool) map[string]interface{} {
	return map[string]interface{}{
		"update_required": required,
	}
}