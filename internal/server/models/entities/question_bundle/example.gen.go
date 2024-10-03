package question_bundle

var ExampleJSON = `{
    bundle_type: BundleType, // How the QuestionBundle is comprised: Manual, ByCategory, or ByTag
    description: string, // Optional further description of the contents of the bundle
    language_id: string, // The default language code for the question bundle
    language_ids: string[], // A list of language id's in which the question bundle is available
    last_export_timestamp: number, // The timestamp of the last export of the question bundle
    question_context_ids: string[], // The list of QuestionContextId strings that are either added manually or queried when the QuestionBundle is saved
    status: Status, // The current status of the question bundle
    subtitle: string, // The subtitle of the question bundle
    title: string, // The title of the question bundle
}`