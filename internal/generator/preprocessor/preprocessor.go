package preprocessor

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bakeable/bkry/internal/generator/types"
	utils "github.com/bakeable/bkry/tools"

	"github.com/jinzhu/inflection"
)


func (p *Preprocessor) GetEntityConfigurations() []types.EntityConfig {
	// Lees alle .json-bestanden in de configuratiemap
	files, err := os.ReadDir(p.ConfigDir)
	if err != nil {
		panic(err)
	}

	entityConfigurations := []types.EntityConfig{}
	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			entity := types.EntityConfig{}
			data, err := os.ReadFile(p.ConfigDir + file.Name())
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(data, &entity)
			if err != nil {
				panic(err)
			}
			entityConfigurations = append(entityConfigurations, entity)
		}
	}

	return entityConfigurations
}

func (p *Preprocessor) GetOverview() types.EntityOverview {
	entityConfigurations := p.GetEntityConfigurations()

	for i, entity := range entityConfigurations {
		entityConfigurations[i].EntityName = utils.PascalCaseToSnakeCase(entity.TypeName)
	}

	p.Mappings = CreateMappings([]GeneralEntity{
		{TypeName: "AuditInfo", EntityName: "auditinfo"},
		{TypeName: "DocRef", EntityName: "docref"},
		{TypeName: "Dimension", EntityName: "dimension"},
	}, entityConfigurations)

	for _, entity := range entityConfigurations {
		utils.SaveAsJSON(p.ConfigDir+entity.TypeName+".json", entity)
	}

	localizedEntities := []types.EntityConfig{}
	for i, entity := range entityConfigurations {
		fmt.Println("\n\nGenerating code for entity: " + entity.TypeName)

		if entity.IsLocalizable {
			localizedEntity := p.PreprocessEntity(entity.DeepCopy(), true)
			localizedEntities = append(localizedEntities, localizedEntity)
		}

		entityConfigurations[i] = p.PreprocessEntity(entity.DeepCopy(), false)
	}

	entityConfigurations = append(entityConfigurations, localizedEntities...)

	overview := types.EntityOverview{
		EntityNameTypes: "\n",
		Entities:        entityConfigurations,
	}

	for _, entity := range entityConfigurations {
		overview.EntityNameTypes += " | '" + entity.TypeName + "'\n"
		overview.EntityNames += " | '" + entity.EntityName + "'\n"
		overview.EntityStoreImports = append(overview.EntityStoreImports, "import { create"+entity.TypeName+"Store } from './entities/"+entity.EntityName+"/store.gen';")
		overview.EntityStoreDefinitions = append(overview.EntityStoreDefinitions, entity.TypeName+": create"+entity.TypeName+"Store,")
	}

	return overview
}

func (p *Preprocessor) PreprocessEntity(entity types.EntityConfig, localized bool) types.EntityConfig {
	if localized {
		entity = p.handleLocalisation(entity)
	}

	p.setEntityNames(&entity)
	p.setParentParameters(&entity)

	fieldParents := []string{}
	entity.Fields = p.processFields(entity.Fields, fieldParents, false)

	if len(entity.Enums) > 0 {
		entity.EnumsExists = true
	}
	p.determineImports(&entity)
	p.mapEnumValuesByType(&entity)
	p.collectChildStructs(&entity)
	p.setCollectionAndDocumentPaths(&entity)

	return entity
}

func (p *Preprocessor) handleLocalisation(entity types.EntityConfig) types.EntityConfig {
	entity.Parents = append(entity.Parents, entity.TypeName)
	entity.Fields = p.removeUnlocalizableFields(entity.Fields)

	parentIdExists := false
	for _, field := range entity.Fields {
		if field.FieldName == entity.TypeName+"ID" {
			parentIdExists = true
		}
	}
	if !parentIdExists {
		entity.Fields = append(entity.Fields, types.FieldConfig{
			FieldName:   entity.TypeName + "ID",
			FieldType:   "string",
			JsonName:    utils.PascalCaseToSnakeCase(entity.TypeName) + "_id",
			Description: "The ID of the parent entity for which this is the child document",
		})
	}

	entity.TypeName += "Localisation"

	return entity
}

func (p *Preprocessor) setEntityNames(entity *types.EntityConfig) {
	entity.EntityName = utils.PascalCaseToSnakeCase(entity.TypeName)
	for _, declarator := range p.ForbiddenDeclarators {
		if entity.EntityName == declarator {
			panic(fmt.Sprintf("The name of entity %s contains a forbidden declarator: %s", entity.TypeName, declarator))
		}
	}

	entity.TypeNameLowercase = strings.ToLower(entity.TypeName[:1]) + entity.TypeName[1:]
	entity.TypeNamePlural = inflection.Plural(entity.TypeName)
	entity.EntityNamePlural = inflection.Plural(entity.EntityName)
	entity.PackagePath = entity.EntityName
	entity.VarName = utils.PascalCaseToCamelCase(entity.TypeName) + "ID"
}

func (p *Preprocessor) setParentParameters(entity *types.EntityConfig) {
	for i, parent := range entity.Parents {
		parentLowercase := strings.ToLower(parent)
		parentCamelCase := utils.PascalCaseToCamelCase(parent)
		parentSnakeCase := utils.PascalCaseToSnakeCase(parent)

		entity.ParentsLowerCase = append(entity.ParentsLowerCase, parentLowercase)
		entity.ParentsCamelCase = append(entity.ParentsCamelCase, parentCamelCase)
		entity.ParentsSnakeCase = append(entity.ParentsSnakeCase, parentSnakeCase)

		entity.CollectionPath += inflection.Plural(parentSnakeCase) + "/[" + parentCamelCase + "ID]/"
		entity.CollectionReference += inflection.Plural(parentSnakeCase) + "/{" + parentCamelCase + "Id}/"

		entity.Endpoint += "/" + parentSnakeCase + "/:" + parentCamelCase + "ID"
		entity.ClientEndpoint += "/" + parentSnakeCase + "/{" + parentSnakeCase + "_id}"

		entity.ParentProperties += "entity." + parent + "ID"
		entity.ParentParams += parentCamelCase + "ID string"
		entity.ParentVars += parentCamelCase + "ID"

		if i < len(entity.Parents)-1 {
			entity.ParentProperties += ", "
			entity.ParentParams += ", "
			entity.ParentVars += ", "
		}

		parentIdExists := false
		for _, field := range entity.Fields {
			if field.FieldName == parent+"ID" {
				parentIdExists = true
			}
		}
		if !parentIdExists {
			entity.Fields = append(entity.Fields, types.FieldConfig{
				FieldName:   parent + "ID",
				FieldType:   "string",
				JsonName:    parentSnakeCase + "_id",
				Description: "The ID of the parent entity for which this is the child document",
			})
		}
	}
}

func (p *Preprocessor) processFields(fields []types.FieldConfig, fieldParents []string, parentIsArray bool) []types.FieldConfig {
	for i, field := range fields {
		fields[i] = p.processField(field, fieldParents, parentIsArray)
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].FieldName < fields[j].FieldName
	})

	return fields
}

func (p *Preprocessor) processField(field types.FieldConfig, fieldParents []string, parentIsArray bool) types.FieldConfig {
	// Verwerken van arrays en standaardwaarden
	if field.IsArray {
		if field.FieldType == "string" {
			field.FieldType = "[]string"
			field.DefaultValue = "[]string{}"
			field.IsArray = false
		}
	}

	// Creëer geformatteerde veldtypen
	field.FieldTypeCapitalized = utils.PascalCaseToCamelCase(field.FieldType)
	field.FieldTypeDecoder = p.Mappings.FieldTypeDecoders[field.FieldType]
	field.FieldTypeIsExternal = strings.Contains(field.FieldType, ".")
	field.FieldTypeClient = p.Mappings.GoTypesToTypeScriptTypes[field.FieldType]
	if field.FieldTypeClient == "" {
		field.FieldTypeClient = field.FieldType
	}

	// Instellen van ouderinformatie voor het veld
	field.Parents = fieldParents
	field.ParentIsArray = parentIsArray
	field.Depth = len(fieldParents)
	field.ParentDepth = len(fieldParents) - 1
	if len(fieldParents) > 0 {
		field.ParentObjectPath = strings.Join(fieldParents, ".") + "."
		field.ParentObjectPathSafe = strings.Join(fieldParents, "?.") + "?."
	} else {
		field.ParentObjectPath = ""
		field.ParentObjectPathSafe = ""
	}

	// Meervoud en enkelvoud van veldnamen
	field.FieldNamePlural = inflection.Plural(field.FieldName)
	field.FieldNameSingular = inflection.Singular(field.FieldName)

	// Standaardwaarden instellen
	if field.FieldType == "string" || field.IsEnum {
		if field.DefaultValue == "" {
			field.DefaultValue = "\"\""
			field.DefaultValueClient = "\"\""
		} else {
			if field.IsEnum {
				field.DefaultValueClient = fmt.Sprintf("%sValues.%s", field.FieldType, field.DefaultValue)
			} else {
				field.DefaultValueClient = fmt.Sprintf("\"%s\"", field.DefaultValue)
			}
		}
	} else if defaultValue, ok := p.Mappings.EmptyDefaultValueMap[field.FieldType]; ok {
		field.DefaultValue = defaultValue
		field.DefaultValueClient = p.Mappings.DefaultValueClientMap[field.FieldType]
	} else {
		field.DefaultValue = ""
		field.DefaultValueClient = ""
	}

	// Recursief verwerken van child fields
	if len(field.Fields) > 0 {
		newFieldParents := append(fieldParents, field.JsonName)
		field.Fields = p.processFields(field.Fields, newFieldParents, field.IsArray)
	}

	return field
}


func (p *Preprocessor) determineImports(entity *types.EntityConfig) {
	// Get field types
	fieldTypes := entity.GetAllFieldTypes()

	for _, fieldType := range fieldTypes {
		if value, ok := p.Mappings.StructImports[fieldType]; ok {
			entity.StructImports = append(entity.StructImports, value)
		}
		if value, ok := p.Mappings.TypeScriptImports[fieldType]; ok {
			entity.TypeScriptImports = append(entity.TypeScriptImports, value)
		}
		if value, ok := p.Mappings.TypeScriptTypeImports[fieldType]; ok {
			entity.TypeScriptTypeImports = append(entity.TypeScriptTypeImports, value)
		}
	}
}

func (p *Preprocessor) mapEnumValuesByType(entity *types.EntityConfig) {
	// Collect all enum values by type
	entity.Enums = entity.GetAllEnumValues()

	for _, enum := range entity.Enums {
		enumExists := false
		for i, enumValueByType := range entity.EnumValuesByType {
			if enumValueByType.EnumType == enum.EnumType {
				entity.EnumValuesByType[i].EnumValues = append(enumValueByType.EnumValues, enum.EnumValue)
				entity.EnumValuesByType[i].EnumValuesString = strings.Join(entity.EnumValuesByType[i].EnumValues, "\" | \"")
				entity.EnumValuesByType[i].EnumValuesFormatted = append(enumValueByType.EnumValuesFormatted, types.EnumValueFormatted{
					Uppercase: strings.ToUpper(enum.EnumValue),
					Lowercase: strings.ToLower(enum.EnumValue),
					Value:     enum.EnumValue,
					SnakeCase: utils.PascalCaseToSnakeCase(enum.EnumValue),
				})
				enumExists = true
			}
		}

		if !enumExists {
			entity.EnumValuesByType = append(entity.EnumValuesByType, types.EnumValueByType{
				EnumType: enum.EnumType,
				EnumValues: []string{enum.EnumValue},
				EnumValuesString: strings.Join([]string{enum.EnumValue}, "\" | \""),
				EnumValuesFormatted: []types.EnumValueFormatted{
					{
						Uppercase: strings.ToUpper(enum.EnumValue),
						Lowercase: strings.ToLower(enum.EnumValue),
						Value:     enum.EnumValue,
						SnakeCase: utils.PascalCaseToSnakeCase(enum.EnumValue),
					},
				},
			})
		}
	}
}

func (p *Preprocessor) collectChildStructs(entity *types.EntityConfig) {
	entity.ChildStructs = entity.GetAllChildStructs()
}

func (p *Preprocessor) setCollectionAndDocumentPaths(entity *types.EntityConfig) {
	entity.CollectionPath += entity.EntityNamePlural
	entity.CollectionName = entity.TypeName
	entity.DocumentPath = entity.CollectionPath + "/[" + entity.VarName + "]"
	entity.Endpoint += "/" + entity.EntityName
	entity.ClientEndpoint += "/" + entity.EntityName
	entity.CollectionReference += "/" + entity.TypeName
}

func (p *Preprocessor) removeUnlocalizableFields(fields []types.FieldConfig) []types.FieldConfig {
	newFields := []types.FieldConfig{}
	for i, field := range fields {
		fields[i].Fields = p.removeUnlocalizableFields(fields[i].Fields)

		if field.IsLocalizable {
			newFields = append(newFields, field)
		}
	}
	return newFields
}
