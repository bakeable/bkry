package media

var ExampleJSON = `{
    content_type: ContentType, // The type of content: Image, Video, or Audio
    description: {
      default_locale: "nl-NL", // The default locale of the description 
      default_value: "", // The default description of the Media file 
      locales: {} satisfies Record<string, string>, // A mapping containing the LanguageCode => Description values 
    }, // The description of the media file
    extension: string, // The file extension
    language_ids: string[], // The array of language codes to which this media applies
    mime_type: string, // The official file mimetype
    multilingual: boolean, // Indicates if the media file is multilingual
    size: string, // The size of the file in Kb
    storage_path: string, // The Firebase Storage location
    url: string, // The public URL to get media directly from the internet
}`