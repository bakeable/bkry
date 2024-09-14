package generator

import (
	"github.com/bakeable/bkry/internal/generator/entities"
)

var configFiles = []string{"config", "struct", "encode", "decode", "enums", "example"}

func configTemplates(entity entities.MetaData) []TemplateFile {
	var templates []TemplateFile
	for _, configFile := range configFiles {
		templates = append(templates, TemplateFile{
			TemplateDir: "config",
			FileName: configFile,
			FileExtension: "go",
			OutputDir: ServerDir + "data/entities/" + entity.PackagePath + "/",
		})
	}
	return templates
}

func GenerateConfig(data []entities.MetaData) {
	for _, entity := range data {
		for _, templateFile := range configTemplates(entity) {
			build(templateFile, entity)
		}
	}
}