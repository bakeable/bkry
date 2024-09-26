package generator

import (
	"github.com/bakeable/bkry/input/entities"
)

var endpoints = []string{"add", "delete", "get", "get_all", "query", "get_all_paginated", "update"}

func controllerTemplates(entity entities.MetaData) []TemplateFile {
	// Collect template file configs
	var templateFiles []TemplateFile

	// Create endpoints for the entity
	for _, endpoint := range endpoints {
		exclude := false
		for _, excludeController := range entity.ExcludedEndpoints {
			if excludeController == endpoint {
				exclude = true
				break
			}
		}

		if !exclude {
			templateFiles = append(templateFiles, TemplateFile{
				TemplateDir: "golang/controllers",
				FileName: endpoint,
				FileExtension: "go",
				OutputDir: ServerDir + "api/" + entity.PackagePath + "/controllers/",
			})
		}
	}

	// Create index file for the entity
	templateFiles = append(templateFiles, TemplateFile{
		TemplateDir: "golang/controllers",
		FileName: "index",
		FileExtension: "go",
		OutputDir: ServerDir + "api/" + entity.PackagePath + "/controllers/",
	})

	return templateFiles
}

func GenerateControllers(data []entities.MetaData) {
	for _, entity := range data {
		for _, templateFile := range controllerTemplates(entity) {
			build(templateFile, entity)
		}
	}
}