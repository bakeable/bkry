package generator

import (
	"github.com/bakeable/bkry/input/entities"
)

type Operation struct {
	Name string
	HasBefore bool
	HasAfter bool
}
var actions = []Operation{
	{"get", false, true},
	{"get_all", false, true},
	{"get_all_paginated", false, true},
	{"save", true, true},
	{"delete", true, true},
	{"query", false, true},
	{"find", false, true},
	{"query_group", false, true},
}


func operationTemplates(entity entities.MetaData) []TemplateFile {
	var templateFiles []TemplateFile
	for _, action := range actions {
		if action.HasBefore {
			templateFiles = append(templateFiles, TemplateFile{
				TemplateDir: "golang/operations",
				FileName: action.Name + ".before",
				FileExtension: "go",
				OutputDir: ServerDir + "data/entities/" + entity.PackagePath + "/operations/",
			})
		}
		if action.HasAfter {
			templateFiles = append(templateFiles, TemplateFile{
				TemplateDir: "golang/operations",
				FileName: action.Name + ".after",
				FileExtension: "go",
				OutputDir: ServerDir + "data/entities/" + entity.PackagePath + "/operations/",
			})
		}

		templateFiles = append(templateFiles, TemplateFile{
			TemplateDir: "golang/operations",
			FileName: action.Name,
			FileExtension: "go",
			OutputDir: ServerDir + "data/entities/" + entity.PackagePath + "/operations/",
			ForceWrite: true,
		})
	}

	return templateFiles
}



func GenerateOperations(data []entities.MetaData) {
	for _, entity := range data {
		for _, templateFile := range operationTemplates(entity) {
			build(templateFile, entity)
		}
	}
}