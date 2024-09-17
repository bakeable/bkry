package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bakeable/bkry/internal/generator/entities"
	utils "github.com/bakeable/bkry/tools"

	"github.com/jinzhu/inflection"
)


func Preprocess(entity entities.MetaData, localized bool) entities.MetaData {
	// Prepare localisation
	if localized {
		entity = handleLocalisation(entity)
	}

	// Entity name is snake_case and type name is PascalCase
	entity.EntityName = utils.PascalCaseToSnakeCase(entity.TypeName)
	for _, declarator := range forbiddenDeclarators {
		if entity.EntityName == declarator {
			panic(fmt.Sprintf("The name of entity %s contains a forbidden declarator: %s", entity.TypeName, declarator))
		}
	}

	// Set first letter of TypeName to lower case
	entity.TypeNameLowercase = strings.ToLower(entity.TypeName[:1]) + entity.TypeName[1:]

	// Set plural type name and entity name
	entity.TypeNamePlural = inflection.Plural(entity.TypeName)
	entity.EntityNamePlural = inflection.Plural(entity.EntityName)

	// Set package path
	entity.PackagePath = entity.EntityName

	// Set varname
	entity.VarName = utils.PascalCaseToCamelCase(entity.TypeName) + "ID"
	
	// Generate the parent parameters
	for i, parent := range entity.Parents {
		parentLowercase := strings.ToLower(parent)
		parentCamelCase := utils.PascalCaseToCamelCase(parent)
		parentSnakeCase := utils.PascalCaseToSnakeCase(parent)
		
		// To lower case
		entity.ParentsLowerCase = append(entity.ParentsLowerCase, parentLowercase)
		entity.ParentsCamelCase = append(entity.ParentsCamelCase, parentCamelCase)
		entity.ParentsSnakeCase = append(entity.ParentsSnakeCase, parentSnakeCase)
		
		// Add parent to collection path
		entity.CollectionPath += parent + "/[" + parentCamelCase + "ID]/"
		entity.CollectionReference += parent + "/{" + parentCamelCase + "Id}/"

		// Add parent to endpoint
		entity.Endpoint += "/" + parentSnakeCase + "/:" + parentCamelCase + "ID"
		entity.ClientEndpoint += "/" + parentSnakeCase + "/{" + parentSnakeCase + "_id}"
		
		// Create input and variables for the parent(s)
		entity.ParentProperties += "entity." + parent + "ID"
		entity.ParentParams += parentCamelCase + "ID string"
		entity.ParentVars += parentCamelCase + "ID"

		// Do not add comma if it's the last parent
		if i < len(entity.Parents) - 1 {
			entity.ParentProperties += ", "
			entity.ParentParams += ", "
			entity.ParentVars += ", "
		}
		
		// Add parent ID to fields if not already present
		parentIdExists := false
		for _, field := range entity.Fields {
			if field.FieldName == parent + "ID" {
				parentIdExists = true
			}
		}
		if !parentIdExists {
			entity.Fields = append(entity.Fields, entities.FieldConfig{
				FieldName: parent + "ID",
				FieldType: "string",
				JsonName: parentSnakeCase + "_id",
				Description: "The ID of the parent entity for which this is the child document",
			})
		}
	}

	// Loop through the fields and preprocess them
	fieldParents := []string{}
	enumTypes := []string{}
	enumConfigs := []entities.EnumConfig{}
	imports := []string{}
	childStructs := []string{}
	entity.Fields, imports, _, entity.Enums, _ = handleFields(entity.Fields, imports, enumTypes, enumConfigs, fieldParents, false, childStructs)

	if len(entity.Enums) > 0 {
		entity.EnumsExists = true
	}

	// Determine the imports
	for _, imp := range imports {
		for key, value := range structImports {
			if imp == key {
				entity.StructImports = append(entity.StructImports, value)
			}
		}
		for key, value := range typeScriptImports {
			if imp == key {
				entity.TypeScriptImports = append(entity.TypeScriptImports, value)
				entity.TypeScriptTypeImports = append(entity.TypeScriptTypeImports, typeScriptTypeImports[key])
			}
		}
	}
	if len(entity.StructImports) > 0 {
		entity.StructImportsNotEmpty = true
	}

	// Map enum values by type
	for _, enum := range entity.Enums {
		enumExists := false
		for i, enumValueByType := range entity.EnumValuesByType {
			if enumValueByType.EnumType == enum.EnumType {
				entity.EnumValuesByType[i].EnumValues = append(enumValueByType.EnumValues, enum.EnumValue)
				entity.EnumValuesByType[i].EnumValuesString = strings.Join(entity.EnumValuesByType[i].EnumValues, "\" | \"")
				entity.EnumValuesByType[i].EnumValuesFormatted = append(enumValueByType.EnumValuesFormatted, entities.EnumValueFormatted{
					Uppercase: strings.ToUpper(enum.EnumValue),
					Lowercase: strings.ToLower(enum.EnumValue),
					Value: enum.EnumValue,
					SnakeCase: utils.PascalCaseToSnakeCase(enum.EnumValue),
				})
				enumExists = true
			}
		}

		if !enumExists {
			entity.EnumValuesByType = append(entity.EnumValuesByType, entities.EnumValueByType{
				EnumType: enum.EnumType,
				EnumValues: []string{enum.EnumValue},
				EnumValuesString: strings.Join([]string{enum.EnumValue}, "\" | \""),
				EnumValuesFormatted: []entities.EnumValueFormatted{
					{
						Uppercase: strings.ToUpper(enum.EnumValue),
						Lowercase: strings.ToLower(enum.EnumValue),
						Value: enum.EnumValue,
						SnakeCase: utils.PascalCaseToSnakeCase(enum.EnumValue),
					},
				},
			})
		}
	}

	// Generate the document and collection path
	entity.CollectionPath += entity.TypeName
	entity.CollectionName = entity.TypeName
	entity.DocumentPath = entity.CollectionPath + "/[" + entity.VarName + "]"
	entity.Endpoint += "/" + entity.EntityName
	entity.ClientEndpoint += "/" + entity.EntityName
	entity.CollectionReference += "/" + entity.TypeName

	return entity
}

func handleFields(fields []entities.FieldConfig, imports []string, enumTypes []string, enumConfigs []entities.EnumConfig, fieldParents []string, parentIsArray bool, childStructs []string) ([]entities.FieldConfig, []string, []string, []entities.EnumConfig, []string) {
	for i, field := range fields {
		// Array string fields
		if field.IsArray {
			if field.FieldType == "string" {
				fields[i].FieldType = "[]string"
				fields[i].DefaultValue = "[]string{}"
				fields[i].IsArray = false
			}
		}
		field = fields[i]

		// Create capitalized field type
		fields[i].FieldTypeCapitalized = utils.PascalCaseToCamelCase(field.FieldType)
		fields[i].FieldTypeDecoder = fieldTypeDecoders[field.FieldType]
		fields[i].FieldTypeIsExternal = strings.Contains(field.FieldType, ".")
		fields[i].FieldTypeClient = goTypesToTypeScriptTypes[field.FieldType]
		if fields[i].FieldTypeClient == "" {
			fields[i].FieldTypeClient = field.FieldType
		}

		// The parents of the field
		fields[i].Parents = fieldParents
		fields[i].ParentIsArray = parentIsArray
		fields[i].ParentObjectPath = strings.Join(fieldParents, ".") + "."
		fields[i].ParentObjectPathSafe = strings.Join(fieldParents, "?.") + "?."
		fields[i].Depth = len(fieldParents)
		fields[i].ParentDepth = len(fieldParents) - 1

		// Make sure only one child struct is created
		if field.IsChildStruct {
			// Check if already in the childStructs slcie
			if utils.Contains(childStructs, field.FieldType) {
				fields[i].ChildStructAlreadyExists = true
			} else {
				childStructs = append(childStructs, field.FieldType)
			}
		}

		// Singular and plural
		fields[i].FieldNamePlural = inflection.Plural(field.FieldName)
		fields[i].FieldNameSingular = inflection.Singular(field.FieldName)

		if field.IsEnum {
			// Check if the enum is already defined
			if utils.Contains(enumTypes, field.FieldType) {
				fields[i].IsDuplicateEnum = true
			} else {
				enumTypes = append(enumTypes, field.FieldType)
				for _, enumValue := range field.EnumValues {
					enumConfigs = append(enumConfigs, entities.EnumConfig{
						EnumType: field.FieldType,
						EnumName: fmt.Sprintf("_%s_%s", field.FieldType, utils.CamelCaseToPascalCase(enumValue)),
						EnumValue: enumValue,
						EnumValueCapitalized: utils.CamelCaseToPascalCase(enumValue),
					})
				}
			}
		}
			
		// Check if any of the structImports keys are in the field types
		for key := range structImports {
			if strings.Contains(field.FieldType, key) {
				if utils.Contains(imports, key) {
					continue
				}
				imports = append(imports, key)
			}
		}
		for key := range typeScriptImports {
			if strings.Contains(field.FieldType, key) {
				if utils.Contains(imports, key) {
					continue
				}
				imports = append(imports, key)
			}
		}

		// Handle default value equal to a string
		if field.FieldType == "[][]int" || field.FieldType == "[]int" || field.FieldType == "[]int8" || field.FieldType == "[]int32" || field.FieldType == "[]int64" || field.FieldType == "[]float64" || field.FieldType == "[]bool" || field.FieldType == "[]string" || field.FieldType == "[]interface{}" || field.FieldType == "[]map[string]interface{}" || field.FieldType == "[]auditinfo.AuditInfo" || field.FieldType == "[]docref.DocRef" || field.FieldType == "dimension.Dimension" || field.FieldType == "[]dimension.Dimension" || field.FieldType == "airline_order.AirlineOrder" || field.FieldType == "[]airline_order.AirlineOrder" || field.FieldType == "airline_pricing.AirlinePricing" || field.FieldType == "[]airline_pricing.AirlinePricing" || field.FieldType == "airline_order_group.AirlineOrderGroup" || field.FieldType == "[]airline_order_group.AirlineOrderGroup" || field.FieldType == "airline_settings.AirlineSettings" || field.FieldType == "examination_task.ExaminationTask" || field.FieldType == "price_layer.PriceLayer" || field.FieldType == "product.Product" || field.FieldType == "[]product.Product" || field.FieldType == "[]price_layer.PriceLayer" || field.FieldType == "[]examination_task.ExaminationTask" || field.FieldType == "[]airline_settings.AirlineSettings" {
			// Handle these complex types
			fields[i].DefaultValueClient = "[] as " + goTypesToTypeScriptTypes[field.FieldType]
		} else if field.FieldType == "map[int]int" || field.FieldType == "map[string]int" || field.FieldType == "map[string]float64" || field.FieldType == "map[string]bool" || field.FieldType == "map[string]string" {
			fields[i].DefaultValueClient = "{} satisfies " + goTypesToTypeScriptTypes[field.FieldType]
		} else if field.FieldType == "string" || field.IsEnum {
			fields[i].DefaultValue = "\"" + field.DefaultValue + "\""
			if field.IsEnum {
				if field.DefaultValue == "" {
					panic(fmt.Sprintf("Enum value %s needs a default value", field.FieldName))
				}
				if field.IsArray {
					fields[i].DefaultValueClient = "[" + field.FieldType + "Values." + field.DefaultValue + "]"
				} else {
					fields[i].DefaultValueClient = field.FieldType + "Values." + field.DefaultValue
				}
			} else {
				fields[i].DefaultValueClient = "\"" + field.DefaultValue + "\""
			}
		} else if field.DefaultValue == "" {
			fields[i].DefaultValue = emptyDefaultValueMap[field.FieldType]
			fields[i].DefaultValueClient = defaultValueClientMap[field.FieldType]
		} else if defaultValue, ok := defaultValueClientMap[field.DefaultValue]; ok {
			fields[i].DefaultValueClient = defaultValue
		} else {
			fields[i].DefaultValueClient = field.DefaultValue
		}

		// Handle child structs
		newFieldParents := append(fieldParents, field.JsonName)
		fields[i].Fields, imports, enumTypes, enumConfigs, childStructs = handleFields(field.Fields, imports, enumTypes, enumConfigs, newFieldParents, field.IsArray, childStructs)
	}

	// Sort the fields alphabetically by FieldName
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].FieldName < fields[j].FieldName
	})
	

	return fields, imports, enumTypes, enumConfigs, childStructs
}


func handleLocalisation(entity entities.MetaData) entities.MetaData {
	// Add parent to localisation
	entity.Parents = append(entity.Parents, entity.TypeName)

	// Remove all fields that are not localizable
	entity.Fields = removeUnlocalizableFields(entity.Fields)

	// Add parent Id to fields if not already present
	parentIdExists := false
	for _, field := range entity.Fields {
		if field.FieldName == entity.TypeName + "ID" {
			parentIdExists = true
		}
	}
	if !parentIdExists {
		entity.Fields = append(entity.Fields, entities.FieldConfig{
			FieldName: entity.TypeName + "ID",
			FieldType: "string",
			JsonName: utils.PascalCaseToSnakeCase(entity.TypeName) + "_id",
			Description: "The ID of the parent entity for which this is the child document",
		})
	}

	// Add localisation to TypeName
	entity.TypeName += "Localisation"

	return entity
}

func removeUnlocalizableFields(fields []entities.FieldConfig) []entities.FieldConfig {
	newFields := []entities.FieldConfig{}
	for i, field := range fields {
		// Remove unlocalizable fields from child structs
		fields[i].Fields = removeUnlocalizableFields(fields[i].Fields)

		if field.IsLocalizable {
			newFields = append(newFields, field)
		}
	}
	return newFields
}