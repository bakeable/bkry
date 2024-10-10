package types

type TemplateFile struct {
	TemplateDir string `json:"template_dir"`
	FileName string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	OutputDir string `json:"output_dir"`
	OutputFileName string `json:"output_file_name"`
	ForceWrite bool `json:"force_write"`
	InitializeOnly bool `json:"initialize_only"`
	InputData string `json:"input_data"`
}


type EntityConfig struct {
	ExcludedEndpoints []string `json:"excludedEndpoints"` // The excluded endpoints for the entity ["get", "get_all", "update", "create", "delete", "query", "find", "query_group"]
	Fields []FieldConfig `json:"fields"` // The fields of the entity
	IsLocalizable bool `json:"isLocalizable"` // Whether the entity is localizable
	Parents []string `json:"parents"` // The parents of the entity
	TypeName string `json:"typeName"` // The singular name of the entity starting with an uppercase letter

	// FILLED BY PREPROCESS
	ClientEndpoint string `json:"clientEndpoint,omitempty"` // Is calculated --- The client endpoints for the entity
	CollectionPath string `json:"collectionPath,omitempty"` // Is calculated --- The path to the collection
	CollectionName string `json:"collectionName,omitempty"` // Is calculated --- The name of the collection
	DocumentPath string `json:"documentPath,omitempty"` // Is calculated --- The path to the document
	CollectionReference string `json:"collectionReference,omitempty"` // Is calculated --- The document reference for the entity to be used on the client side
	ChildStructs []FieldConfig `json:"childStructs,omitempty"` // Is calculated --- The child structs of the entity
	Endpoint string `json:"endpoint,omitempty"` // Is calculated --- The endpoint for the entity
	EntityName string `json:"entityName,omitempty"` // Is calculated --- The name of the entity
	EntityNamePlural string `json:"entityNamePlural,omitempty"` // Is calculated --- The plural name of the entity
	Enums []EnumValue `json:"enums,omitempty"` // Is calculated --- The enums of the entity
	EnumValuesByType []EnumValueByType `json:"enumValuesByType,omitempty"` // Is calculated --- The enum values of the entity by enum type
	EnumsExists bool `json:"enumsExists,omitempty"` // Is calculated --- Whether the entity has enums
	TypeNamePlural string `json:"typeNamePlural,omitempty"` // Is calculated --- The plural name of the entity
	TypeNameLowercase string `json:"typeNameLowercase,omitempty"` // Is calculated --- The name of the entity starting with a lowercase letter
	TypeScriptImports []string `json:"typeScriptImports,omitempty"` // Is calculated --- The imports for building the TypeScript interface
	TypeScriptTypeImports []string `json:"typeScriptTypeImports,omitempty"` // Is calculated --- The imports for building the TypeScript type
	ParentsLowerCase []string `json:"parentsLowerCase,omitempty"` // Is calculated --- The parents of the entity in lower case
	ParentsCamelCase []string `json:"parentsCamelCase,omitempty"` // Is calculated --- The parents of the entity in camel case
	ParentsPascalCase []string `json:"parentsPascalCase,omitempty"` // Is calculated --- The parents of the entity in pascal case
	ParentsSnakeCase []string `json:"parentsSnakeCase,omitempty"` // Is calculated --- The parents of the entity in snake case
	ParentParams string `json:"parentParams,omitempty"` // Is calculated --- The parameters for the parent
	ParentProperties string `json:"parentProperties,omitempty"` // Is calculated --- The properties for the parent
	ParentVars string `json:"parentVars,omitempty"` // Is calculated --- The variables for the parent
	PackagePath string `json:"packagePath,omitempty"` // Is calculated --- The path to the package
	StructImports []string `json:"structImports,omitempty"` // Is calculated --- The imports for building the struct
	StructImportsNotEmpty bool `json:"structImportsNotEmpty,omitempty"` // Is calculated --- Whether the struct imports exist and is not empty
	VarName string `json:"varName,omitempty"` // Is calculated --- The variable name of the Id of the entity
}

type FieldConfig struct {
	FieldName string `json:"fieldName"` // The name of the field
	FieldType string `json:"fieldType"` // The type of the field
	IsArray bool `json:"isArray"` // Whether the field is an array
	IsChildStruct bool `json:"isChildStruct"` // Whether the field is a child struct
	IsEnum bool `json:"isEnum"` // Whether the field is an enum
	IsLocalizable bool `json:"isLocalizable"` // Whether the field is localizable
	EnumValues []string `json:"enumValues,omitempty"` // The values of the enum
	JsonName string `json:"jsonName"` // The name of the field in the JSON
	DefaultValue string `json:"defaultValue"` // The default value of the field
	Fields []FieldConfig `json:"fields"` // The fields of the field
	Description string `json:"description"` // The description of the field
	IsRequired bool `json:"isRequired"` // Whether the field is required
	IsFixedAcrossLocalisations bool `json:"isFixedAcrossLocalisations"` // Whether the field is fixed across localisations
	ArrayMatchOn string `json:"arrayMatchOn"` // The field to match on when the field is an array, empty if matching is done order-wise

	// FILLED BY PREPROCESS
	DefaultValueClient string `json:"defaultValueClient,omitempty"` // Is calculated --- The default value for the client side
	FieldNameSingular string `json:"fieldNameSingular,omitempty"` // Is calculated --- The singular name of the field
	FieldNamePlural string `json:"fieldNamePlural,omitempty"` // Is calculated --- The plural name of the field
	FieldTypeCapitalized string `json:"fieldTypeCapitalized,omitempty"` // Is calculated --- The name of the field type starting with an uppercase letter
	FieldTypeDecoder string `json:"fieldTypeDecoder,omitempty"` // Is calculated --- The decoder function name for the field type
	FieldTypeIsExternal bool `json:"fieldTypeIsExternal,omitempty"` // Is calculated --- Whether the field type is external
	FieldTypeClient string `json:"fieldTypeClient,omitempty"` // Is calculated --- The client field type
	IsDuplicateEnum bool `json:"isDuplicateEnum,omitempty"` // Is calculated --- Whether the field is a duplicate enum, i.e. the enum is already defined
	Parents []string `json:"parents,omitempty"` // Is calculated --- The parents of the field
	ParentIsArray bool `json:"parentIsArray,omitempty"` // Is calculated --- Whether the parent is an array
	ParentObjectPath string `json:"parentObjectPath,omitempty"` // Is calculated --- The path to the parent JSON object, separated by dots
	ParentObjectPathSafe string `json:"parentObjectPathSafe,omitempty"` // Is calculated --- The path to the parent JSON object, separated by dots, with safe ? characters
	Depth int `json:"depth,omitempty"` // Is calculated --- The depth of the field, i.e. how many parents the field has
	ParentDepth int `json:"parentDepth,omitempty"` // Is calculated --- The depth of the parent, i.e. how many parents the parent has
	EmbeddedChildStructs []string `json:"embeddedChildStructs,omitempty"` // Is calculated --- The list of unique embedded child structs
	ChildStructAlreadyExists bool `json:"childStructAlreadyExists,omitempty"` // Is calculated --- Whether the child struct already exists
}

type EnumValue struct {
	EnumType string `json:"enumType,omitempty"` // Is calculated --- The type of the enum
	EnumValue string `json:"enumValue,omitempty"` // Is calculated --- The values of the enum
	EnumName string `json:"enumName,omitempty"` // Is calculated --- The name of the enum
	EnumValueCapitalized string `json:"enumValueCapitalized,omitempty"` // Is calculated --- The values of the enum starting with an uppercase letter
}

type EnumValueByType struct {
	EnumType string `json:"enumType,omitempty"` // The type of the enum
	EnumValues []string `json:"enumValues,omitempty"` // The values of the enum
	EnumValuesString string `json:"enumValuesString,omitempty"` // The values of the enum as a string, separated by |
	EnumValuesFormatted []EnumValueFormatted `json:"enumValuesFormatted,omitempty"` // The formatted enum values
}

type EnumValueFormatted struct {
	Value string `json:"value,omitempty"` // The enum value
	Uppercase string `json:"uppercase,omitempty"` // The uppercase enum values
	Lowercase string `json:"lowercase,omitempty"` // The lowercase enum values
	SnakeCase string `json:"snakeCase,omitempty"` // The snake case enum values
}

type EntityOverview struct {
	EntityNameTypes string `json:"entityNameTypes,omitempty"`
	EntityStoreImports []string `json:"entityStoreImports,omitempty"`
	EntityStoreDefinitions []string `json:"entityStoreDefinitions,omitempty"`
	EntityNames string `json:"entityNames,omitempty"`
	Entities []EntityConfig `json:"entities,omitempty"`
}
