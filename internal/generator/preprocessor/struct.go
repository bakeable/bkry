package preprocessor

type Preprocessor struct {
	ConfigDir              string
	Mappings               Mappings
	ForbiddenDeclarators   []string
}

type GeneralEntity struct {
	TypeName string
	EntityName string
}

type Mappings struct {
	StructImports map[string]string
	TypeScriptImports map[string]string
	TypeScriptTypeImports map[string]string
	FieldTypeDecoders map[string]string
	EmptyDefaultValueMap map[string]string
	GoTypesToTypeScriptTypes map[string]string
	DefaultValueClientMap map[string]string
}