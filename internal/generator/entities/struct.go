package entities

type MetaData struct {
	// TO BE FILLED
	ExcludedEndpoints []string // The excluded endpoints for the entity ["get", "get_all", "update", "create", "delete", "query", "find", "query_group"]
	Fields []FieldConfig // The fields of the entity
	IsLocalizable bool // Whether the entity is localizable
	Parents []string // The parents of the entity
	TypeName    string // The singular name of the entity starting with an uppercase letter


	// FILLED BY PREPROCESS
	ClientEndpoint string // Is calculated --- The client endpoints for the entity
	CollectionPath string // Is calculated --- The path to the collection
	CollectionName string // Is calculated --- The name of the collection
	DocumentPath string // Is calculated --- The path to the document
	CollectionReference string // Is calculated --- The document reference for the entity to be used on the client side
	Endpoint string // Is calculated --- The endpoint for the entity
	EntityName  string // Is calculated --- The name of the entity
	EntityNamePlural string // Is calculated --- The plural name of the entity
	Enums []EnumConfig // Is calculated --- The enums of the entity
	EnumValuesByType []EnumValueByType // Is calculated --- The enum values of the entity by enum type
	EnumsExists bool // Is calculated --- Whether the entity has enums
	TypeNamePlural string // Is calculated --- The plural name of the entity
	TypeNameLowercase string // Is calculated --- The name of the entity starting with a lowercase letter
	TypeScriptImports []string // Is calculated --- The imports for building the TypeScript interface
	TypeScriptTypeImports []string // Is calculated --- The imports for building the TypeScript type
	ParentsLowerCase []string // Is calculated --- The parents of the entity in lower case
	ParentsCamelCase []string // Is calculated --- The parents of the entity in camel case
	ParentsPascalCase []string // Is calculated --- The parents of the entity in pascal case
	ParentsSnakeCase []string // Is calculated --- The parents of the entity in snake case
	ParentParams string // Is calculated --- The parameters for the parent
	ParentProperties string // Is calculated --- The properties for the parent
	ParentVars string // Is calculated --- The variables for the parent
	PackagePath string // Is calculated --- The path to the package
	StructImports []string // Is calculated --- The imports for building the struct
	StructImportsNotEmpty bool // Is calculated --- Whether the struct imports exist and is not empty
	VarName     string // Is calculated --- The variable name of the Id of the entity
}

func (m *MetaData) DeepCopy() MetaData {
	copy := MetaData{
		ExcludedEndpoints: make([]string, len(m.ExcludedEndpoints)),
		Fields: make([]FieldConfig, len(m.Fields)),
		IsLocalizable: m.IsLocalizable,
		Parents: make([]string, len(m.Parents)),
		TypeName: m.TypeName,
	}
	for i, v := range m.ExcludedEndpoints {
		copy.ExcludedEndpoints[i] = v
	}
	for i, v := range m.Parents {
		copy.Parents[i] = v
	}
	for i, v := range m.Fields {
		copy.Fields[i] = v.DeepCopy()
	}
	return copy
}

type FieldConfig struct {
	// TO BE FILLED
	FieldName string // The name of the field
	FieldType string // The type of the field
	IsArray bool // Whether the field is an array
	IsChildStruct bool // Whether the field is a child struct
	IsEnum bool // Whether the field is an enum
	IsLocalizable bool // Whether the field is localizable
	EnumValues []string // The values of the enum
	JsonName string // The name of the field in the JSON
	DefaultValue string // The default value of the field
	Fields []FieldConfig // The fields of the field
	Description string // The description of the field
	IsRequired bool // Whether the field is required
	IsFixedAcrossLocalisations bool // Whether the field is fixed across localisations
	ArrayMatchOn string // The field to match on when the field is an array, empty if matching is done order-wise
	

	// FILLED BY PREPROCESS
	DefaultValueClient string // Is calculated --- The default value for the client side
	FieldNameSingular string // Is calculated --- The singular name of the field
	FieldNamePlural string // Is calculated --- The plural name of the field
	FieldTypeCapitalized string // Is calculated --- The name of the field type starting with an uppercase letter
	FieldTypeDecoder string // Is calculated --- The decoder function name for the field type
	FieldTypeIsExternal bool // Is calculated --- Whether the field type is external
	FieldTypeClient string // Is calculated --- The client field type
	IsDuplicateEnum bool // Is calculated --- Whether the field is a duplicate enum, i.e. the enum is already defined
	Parents []string // Is calculated --- The parents of the field
	ParentIsArray bool // Is calculated --- Whether the parent is an array
	ParentObjectPath string // Is calculated --- The path to the parent JSON object, separated by dots
	ParentObjectPathSafe string // Is calculated --- The path to the parent JSON object, separated by dots, with safe ? characters
	Depth int // Is calculated --- The depth of the field, i.e. how many parents the field has
	ParentDepth int // Is calculated --- The depth of the parent, i.e. how many parents the parent has
	EmbeddedChildStructs []string // Is calculated --- The list of unique embedded child structs
	ChildStructAlreadyExists bool // Is calculated --- Whether the child struct already exists
}

func (f *FieldConfig) DeepCopy() FieldConfig {
	copy := FieldConfig{
		FieldName: f.FieldName,
		FieldType: f.FieldType,
		IsArray: f.IsArray,
		IsChildStruct: f.IsChildStruct,
		IsEnum: f.IsEnum,
		IsLocalizable: f.IsLocalizable,
		JsonName: f.JsonName,
		DefaultValue: f.DefaultValue,
		EnumValues: make([]string, len(f.EnumValues)),
		Fields: make([]FieldConfig, len(f.Fields)),
		Description: f.Description,
		IsFixedAcrossLocalisations: f.IsFixedAcrossLocalisations,
		ArrayMatchOn: f.ArrayMatchOn,
	}
	for i, v := range f.EnumValues {
		copy.EnumValues[i] = v
	}
	for i, v := range f.Fields {
		copy.Fields[i] = v.DeepCopy()
	}
	return copy
}

type EnumConfig struct {
	EnumType string // Is calculated --- The type of the enum
	EnumValue string // Is calculated --- The values of the enum
	EnumName string // Is calculated --- The name of the enum
	EnumValueCapitalized string // Is calculated --- The values of the enum starting with an uppercase letter
}

type EnumValueByType struct {
	EnumType string // The type of the enum
	EnumValues []string // The values of the enum
	EnumValuesString string // The values of the enum as a string, separated by |
	EnumValuesFormatted []EnumValueFormatted // The formatted enum values
}

type EnumValueFormatted struct {
	Value string // The enum value
	Uppercase string // The uppercase enum values
	Lowercase string // The lowercase enum values
	SnakeCase string // The snake case enum values
}