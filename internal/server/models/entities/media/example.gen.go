package media

var ExampleJSON = `{
    contentType: ContentType, // The type of content: image, video, audio or file
    description: string, // The description of the media
    extension: string, // The file extension
    filename: string, // The filename of the media
    mimeType: string, // The official file mimetype
    size: string, // The size of the file in Kb
    storagePath: string, // The Firebase Storage location
    url: string, // The public URL to get media directly from the internet
}`