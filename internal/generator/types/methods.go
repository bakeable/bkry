package types

import (
	"encoding/json"
	"fmt"

	utils "github.com/bakeable/bkry/tools"
)


func (m *EntityConfig) GetAllFieldTypes() []string {
	types := make([]string, 0)
	for _, f := range m.Fields {
		types = append(types, f.FieldType)
		types = append(types, f.GetAllFieldTypes()...)
	}
	fmt.Println("types", types)
	return utils.Unique(types)
}

func (m *FieldConfig) GetAllFieldTypes() []string {
	types := make([]string, 0)
	for _, f := range m.Fields {
		types = append(types, f.FieldType)
		types = append(types, f.GetAllFieldTypes()...)
	}
	return types
}

func (m *EntityConfig) GetAllEnumTypes() []string {
	types := make([]string, 0)
	for _, f := range m.Fields {
		if f.IsEnum {
			types = append(types, f.FieldType)
		}
		types = append(types, f.GetAllEnumTypes()...)
	}
	return utils.Unique(types)
}

func (m *FieldConfig) GetAllEnumTypes() []string {
	types := make([]string, 0)
	for _, f := range m.Fields {
		if f.IsEnum {
			types = append(types, f.FieldType)
		}
		types = append(types, f.GetAllEnumTypes()...)
	}
	return types
}

func (m *EntityConfig) GetAllEnumValues() []EnumValue {
	enumValues := make([]EnumValue, 0)
	for _, f := range m.Fields {
		if f.IsEnum {
			for _, enumValue := range f.EnumValues {
				enumValues = append(enumValues, EnumValue{
					EnumType:             f.FieldType,
					EnumName:             fmt.Sprintf("%s_%s", f.FieldType, utils.CamelCaseToPascalCase(enumValue)),
					EnumValue:            enumValue,
					EnumValueCapitalized: utils.CamelCaseToPascalCase(enumValue),
				})
			}
		}
		enumValues = append(enumValues, f.GetAllEnumValues()...)
	}

	// Get unique by EnumName
	uniqueEnumValues := make([]EnumValue, 0)
	uniqueEnumValuesMap := make(map[string]bool)
	for _, enumValue := range enumValues {
		if _, ok := uniqueEnumValuesMap[enumValue.EnumName]; !ok {
			uniqueEnumValuesMap[enumValue.EnumName] = true
			uniqueEnumValues = append(uniqueEnumValues, enumValue)
		}
	}

	return uniqueEnumValues
}

func (m *FieldConfig) GetAllEnumValues() []EnumValue {
	enumValues := make([]EnumValue, 0)
	if m.IsEnum {
		for _, enumValue := range m.EnumValues {
			enumValues = append(enumValues, EnumValue{
				EnumType:             m.FieldType,
				EnumName:             fmt.Sprintf("%s_%s", m.FieldType, utils.CamelCaseToPascalCase(enumValue)),
				EnumValue:            enumValue,
				EnumValueCapitalized: utils.CamelCaseToPascalCase(enumValue),
			})
		}
	}
	for _, f := range m.Fields {
		enumValues = append(enumValues, f.GetAllEnumValues()...)
	}
	return enumValues
}


func (m *EntityConfig) GetAllChildStructs() []FieldConfig {
	childStructs := make([]FieldConfig, 0)
	for _, f := range m.Fields {
		if f.IsChildStruct {
			childStructs = append(childStructs, f)
		}
		childStructs = append(childStructs, f.GetAllChildStructs()...)
	}
	
	// Get unique by FieldType
	uniqueChildStructs := make([]FieldConfig, 0)
	uniqueChildStructsMap := make(map[string]bool)
	for _, childStruct := range childStructs {
		if _, ok := uniqueChildStructsMap[childStruct.FieldType]; !ok {
			uniqueChildStructsMap[childStruct.FieldType] = true
			uniqueChildStructs = append(uniqueChildStructs, childStruct)
		}
	}

	return uniqueChildStructs
}


func (m *FieldConfig) GetAllChildStructs() []FieldConfig {
	childStructs := make([]FieldConfig, 0)
	for _, f := range m.Fields {
		if f.IsChildStruct {
			childStructs = append(childStructs, f)
		}
		childStructs = append(childStructs, f.GetAllChildStructs()...)
	}

	return childStructs
}


func (m *EntityConfig) DeepCopy() EntityConfig {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	var copy EntityConfig
	err = json.Unmarshal(data, &copy)
	if err != nil {
		panic(err)
	}
	return copy
}



func (f *FieldConfig) DeepCopy() FieldConfig {
	data, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	var copy FieldConfig
	err = json.Unmarshal(data, &copy)
	if err != nil {
		panic(err)
	}
	return copy
}