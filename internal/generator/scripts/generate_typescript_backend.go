package generator

import (
	"path/filepath"

	"github.com/bakeable/bkry/internal/generator/entities"
)

var BackEndDir = OutputDir + "functions"
var backEndDataDir = BackEndDir + "/src/data/"
var backEndTemplateDir = "typescript/back-end/entities/"
var backEndOutputDir = backEndDataDir + "entities/"
var backEndTemplates = []TemplateFile{
	{
		FileName:      "dto",
		FileExtension: "d.ts",
	},
	{
		FileName:      "dto",
		FileExtension: "ts",
	},
	{
		FileName:      "entity",
		FileExtension: "d.ts",
		InitializeOnly: true,
	},
	{
		FileName:      "entity",
		FileExtension: "ts",
		InitializeOnly: true,
	},
	{
		FileName:      "index",
		FileExtension: "ts",
		InitializeOnly: true,
	},
	{
		FileName:      "list",
		FileExtension: "d.ts",
	},
	{
		FileName:      "list",
		FileExtension: "ts",
	},
	{
		FileName:      "parsers",
		FileExtension: "ts",
	},
	{
		FileName:      "custom_validators",
		FileExtension: "ts",
		InitializeOnly: true,
	},
	{
		FileName:      "validators",
		FileExtension: "ts",
	},
}

func GenerateBackend(data []entities.MetaData) {
	for _, entity := range data {
		generateBackendEntity(entity)
	}

	generateBackendEntityOverview(data)
}

func generateBackendEntity(entity entities.MetaData) {
	for _, templateFile := range backEndTemplates {
		// Fill in the template file
		templateFile.TemplateDir = backEndTemplateDir
		templateFile.OutputDir = backEndOutputDir + entity.EntityName + "/"
		
		build(templateFile, entity)
	}
}


var overviewTemplates = []TemplateFile{
	{
		TemplateDir: "typescript/back-end",
		FileName: "types",
		FileExtension: "d.ts",
		OutputDir: filepath.Join(backEndDataDir, "types/"),
		ForceWrite: true,
	},
	{
		TemplateDir: "typescript/back-end",
		FileName: "types",
		FileExtension: "ts",
		OutputDir: filepath.Join(backEndDataDir, "types/"),
		ForceWrite: true,
	},
	{
		TemplateDir: "typescript/back-end",
		FileName: "index",
		FileExtension: "ts",
		OutputDir: filepath.Join(backEndDataDir, "entities/"),
		OutputFileName: "index.ts",
		ForceWrite: true,
	},
}
type ClientOverview struct {
	EntityNameTypes string
	EntityStoreImports []string
	EntityStoreDefinitions []string
	EntityNames string
	Entities []entities.MetaData
}
func generateBackendEntityOverview(data []entities.MetaData) {
	// Build the type
	overview := ClientOverview{
		EntityNameTypes: "\n",
		Entities: data,
	}

	for _, entity := range data {
		overview.EntityNameTypes += " | '" + entity.TypeName + "'\n"
		overview.EntityNames += " | '" + entity.EntityName + "'\n"
		overview.EntityStoreImports = append(overview.EntityStoreImports, "import { create" + entity.TypeName + "Store } from './entities/" + entity.EntityName + "/store.gen';")
		overview.EntityStoreDefinitions = append(overview.EntityStoreDefinitions, entity.TypeName + ": create" + entity.TypeName + "Store,")
	}

	for _, templateFile := range overviewTemplates {
		build(templateFile, overview)
	}
}
