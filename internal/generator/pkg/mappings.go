package generator

type GeneralEntity struct {
	TypeName string
	EntityName string
}

func CreateMappings(general []GeneralEntity, entities []EntityConfig) {
	for _, entity := range entities {
		structImports[entity.TypeName + "." + entity.EntityName] = "github.com/bakeable/bkry/internal/server/models/entities/" + entity.TypeName
		typeScriptImports[entity.TypeName + "." + entity.EntityName] = "import { " + entity.EntityName + " } from '../" + entity.TypeName + "'"
		typeScriptTypeImports[entity.TypeName + "." + entity.EntityName] = "import { type " + entity.EntityName + " } from '../" + entity.TypeName + "'"
		fieldTypeDecoders[entity.TypeName + "." + entity.EntityName] = entity.TypeName + ".Decode"
		fieldTypeDecoders["[]" + entity.TypeName + "." + entity.EntityName] = entity.TypeName + ".DecodeAll"
		emptyDefaultValueMap[entity.TypeName + "." + entity.EntityName] = entity.TypeName + "." + entity.EntityName + "{}"
		emptyDefaultValueMap["[]" + entity.TypeName + "." + entity.EntityName] = "[]" + entity.TypeName + "." + entity.EntityName + "{}"
		goTypesToTypeScriptTypes[entity.TypeName + "." + entity.EntityName] = entity.EntityName
		defaultValueClientMap["[]" + entity.TypeName + "." + entity.EntityName] = "[] as " + entity.EntityName + "[]"
	}

	// General types
	for _, entity := range general {
		structImports[entity.TypeName + "." + entity.EntityName] = "github.com/bakeable/bkry/internal/server/models/general/" + entity.TypeName
		typeScriptImports[entity.TypeName + "." + entity.EntityName] = "import { " + entity.EntityName + " } from '../" + entity.TypeName + "'"
		typeScriptTypeImports[entity.TypeName + "." + entity.EntityName] = "import { type " + entity.EntityName + " } from '../" + entity.TypeName + "'"
		fieldTypeDecoders[entity.TypeName + "." + entity.EntityName] = entity.TypeName + ".Decode"
		fieldTypeDecoders["[]" + entity.TypeName + "." + entity.EntityName] = entity.TypeName + ".DecodeAll"
		emptyDefaultValueMap[entity.TypeName + "." + entity.EntityName] = entity.TypeName + "." + entity.EntityName + "{}"
		emptyDefaultValueMap["[]" + entity.TypeName + "." + entity.EntityName] = "[]" + entity.TypeName + "." + entity.EntityName + "{}"
		goTypesToTypeScriptTypes[entity.TypeName + "." + entity.EntityName] = entity.EntityName
		defaultValueClientMap["[]" + entity.TypeName + "." + entity.EntityName] = "[] as " + entity.EntityName + "[]"
	}
}

var structImports = map[string]string{}
var typeScriptImports = map[string]string{}
var typeScriptTypeImports = map[string]string{}
var fieldTypeDecoders = map[string]string{
	"time.Time": "utils.DecodeTime",
	"[]time.Time": "utils.DecodeTimeArray",
	"string": "utils.DecodeString",
	"int": "utils.DecodeInt",
	"int8": "utils.DecodeInt8",
	"int32": "utils.DecodeInt32",
	"int64": "utils.DecodeInt64",
	"float32": "utils.DecodeFloat32",
	"float64": "utils.DecodeFloat64",
	"bool": "utils.DecodeBool",
	"[]string": "utils.DecodeStringArray",
	"[]int": "utils.DecodeIntArray",
	"[][]int": "utils.DecodeIntArrayArray",
	"[]int8": "utils.DecodeInt8Array",
	"[]int32": "utils.DecodeInt32Array",
	"[]int64": "utils.DecodeInt64Array",
	"[]float64": "utils.DecodeFloatArray",
	"[]bool": "utils.DecodeBoolArray",
	"[]interface{}": "utils.DecodeInterfaceArray",
	"[]map[string]interface{}": "utils.DecodeMapArray",
	"interface{}": "utils.DecodeInterface",
	"map[string]interface{}": "utils.DecodeMap",
	"map[string]int": "utils.DecodeIntMap",
	"map[string]float64": "utils.DecodeFloatMap",
	"map[string]bool": "utils.DecodeBoolMap",
	"map[string]string": "utils.DecodeStringMap",
	"map[int]int": "utils.DecodeIntIntMap",
}
var emptyDefaultValueMap = map[string]string{
	"string": "\"\"",
	"int": "0",
	"int8": "0",
	"int32": "0",
	"int64": "0",
	"float64": "0",
	"bool": "false",
	"interface{}": "nil",
	"map[string]interface{}": "nil",
	"map[string]int": "map[string]int{}",
	"map[string]float64": "map[string]float64{}",
	"map[string]bool": "map[string]bool{}",
	"map[string]string": "map[string]string{}",
	"[]map[string]interface{}": "[]map[string]interface{}{}",
	"[]map[string]int": "[]map[string]int{}",
	"[]map[string]float64": "[]map[string]float64{}",
	"[]map[string]bool": "[]map[string]bool{}",
	"[]map[string]string": "[]map[string]string{}",
	"[]string": "[]string{}",
	"[]int": "[]int{}",
	"[]float64": "[]float64{}",
	"[]bool": "[]bool{}",
	"[]interface{}": "[]interface{}{}",
}
var forbiddenDeclarators = []string{
	"package",
	"var",
	"import",
	"const",
	"func",
	"return",
	"break",
	"continue",
	"goto",
	"fallthrough",
	"defer",
	"if",
	"else",
	"switch",
	"case",
	"default",
	"select",
	"for",
	"range",
	"type",
	"struct",
	"interface",
	"chan",
	"map",
}
var goTypesToTypeScriptTypes = map[string]string{
    "string":       "string",
    "bool":         "boolean",
    "int":          "number",
    "int8":         "number",
    "int16":        "number",
    "int32":        "number",
    "int64":        "number",
    "uint":         "number",
    "uint8":        "number",
    "uint16":       "number",
    "uint32":       "number",
    "uint64":       "number",
    "uintptr":      "number",
    "float32":      "number",
    "float64":      "number",
    "complex64":    "any",
    "complex128":   "any",
    "byte":         "number",
    "rune":         "string",
    "map[string]string": "Record<string, string>",
    "map[string]interface{}": "Record<string, any>",
	"map[string]int": "Record<string, number>",
	"map[string]float64": "Record<string, number>",
	"map[string]bool": "Record<string, boolean>",
    "[]string":     "string[]",
    "[]bool":       "boolean[]",
    "[]int":        "number[]",
    "[]int8":       "number[]",
    "[]int16":      "number[]",
    "[]int32":      "number[]",
    "[]int64":      "number[]",
    "[]uint":       "number[]",
    "[]uint8":      "number[]",
    "[]uint16":     "number[]",
    "[]uint32":     "number[]",
    "[]uint64":     "number[]",
    "[]uintptr":    "number[]",
    "[]float32":    "number[]",
    "[]float64":    "number[]",
    "[]complex64":  "any[]",
    "[]complex128": "any[]",
    "[]byte":       "number[]",
    "[]rune":       "string[]",
	"[]interface{}": "any[]",
	"[]map[string]interface{}": "Record<string, any>[]",
	"[][]int": "number[][]",
	"map[int]int": "Record<number, number>",
}
var defaultValueClientMap = map[string]string{
	"string": "\"\"",
	"int": "0",
	"int8": "0",
	"int32": "0",
	"int64": "0",
	"float64": "0",
	"bool": "false",
	"interface{}": "nil",
	"map[string]interface{}": "{} satisfies Record<string, any>",
	"map[string]int": "{} satisfies Record<string, number>",
	"map[string]float64": "{} satisfies Record<string, number>",
	"map[string]bool": "{} satisfies Record<string, boolean>",
	"map[string]string": "{} satisfies Record<string, string>",
	"[]string": "[] as string[]",
	"[]int": "[] as number[]",
	"[][]int": "[] as Array<number[]>",
	"[]float64": "[] as number[]",
	"[]bool": "[] as boolean[]",
	"[]interface{}": "[]",
	"[]map[string]interface{}": "[]",
	"[]string{}": "[] as string[]",
	"[]int{}": "[] as number[]",
	"[]float64{}": "[] as number[]",
	"[]bool{}": "[] as boolean[]",
	"[]interface{}{}": "[]",
	"[]map[string]interface{}{}": "[]",
}