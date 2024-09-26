package entities

var mediaFields = []FieldConfig{
	{
		FieldName: "Extension",
		FieldType: "string",
		JsonName: "extension",
		Description: "The file extension",
	},
	{
		FieldName: "ContentType",
		FieldType: "ContentType",
		IsEnum: true,
		DefaultValue: "Image",
		EnumValues: []string{"image", "video", "audio", "file"},
		JsonName: "contentType",
		Description: "The type of content: image, video, audio or file",
	},
	{
		FieldName: "Description",
		FieldType: "string",
		JsonName: "description",
		Description: "The description of the media",
	},
	{
		FieldName: "MimeType",
		FieldType: "string",
		JsonName: "mimeType",
		Description: "The official file mimetype",
	},
	{
		FieldName: "Size",
		FieldType: "string",
		JsonName: "size",
		Description: "The size of the file in Kb",
	},
	{
		FieldName: "StoragePath",
		FieldType: "string",
		JsonName: "storagePath",
		Description: "The Firebase Storage location",
	},
	{
		FieldName: "Url",
		FieldType: "string",
		JsonName: "url",
		Description: "The public URL to get media directly from the internet",
	},
	{
		FieldName: "Filename",
		FieldType: "string",
		JsonName: "filename",
		Description: "The filename of the media",
	},
}
