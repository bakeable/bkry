package product_package

var ExampleJSON = `{
    description: string, // Optional further description of the contents of the package
    items: {
      entity_type: EntityTypeValues.QuestionBundle, // The entity type that is sold as an item of the package (an enum with possible values: [QuestionBundle GameModus])
      id: "", // The Id of the entity being sold 
    }[], // The items included in the package
    language_id: string, // The default language code for the package
    language_ids: string[], // A list of language id's in which the question bundle is available
    subtitle: string, // The subtitle of the package
    title: string, // The title of the package
}`