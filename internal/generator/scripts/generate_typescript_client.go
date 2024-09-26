package generator

import (
	"path/filepath"

	"github.com/bakeable/bkry/input/entities"
)
var ClientDir = OutputDir + "client"
var clientDataDir = ClientDir + "/src/data/"
var clientTemplateDir = "typescript/client/entities/"
var clientOutputDir = clientDataDir + "entities/"
var clientTemplates = []TemplateFile{
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
		InitializeOnly: true,
	},
	{
		FileName:      "store",
		FileExtension: "d.ts",
	},
	{
		FileName:      "store",
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

func GenerateClient(data []entities.MetaData) {
	for _, entity := range data {
		generateClientEntity(entity)
	}

	generateClientEntityOverview(data)
}

func generateClientEntity(entity entities.MetaData) {
	for _, templateFile := range clientTemplates {
		// Fill in the template file
		templateFile.TemplateDir = clientTemplateDir
		templateFile.OutputDir = clientOutputDir + entity.EntityName + "/"
		
		build(templateFile, entity)
	}
}


var clientOverviewTemplates = []TemplateFile{
	{
		TemplateDir: "typescript/client",
		FileName: "stores",
		FileExtension: "ts",
		OutputDir: clientDataDir,
	},
	{
		TemplateDir: "typescript/client",
		FileName: "types",
		FileExtension: "d.ts",
		OutputDir: filepath.Join(clientDataDir, "types/"),
		ForceWrite: true,
	},
	{
		TemplateDir: "typescript/client",
		FileName: "types",
		FileExtension: "ts",
		OutputDir: filepath.Join(clientDataDir, "types/"),
		ForceWrite: true,
	},
	{
		TemplateDir: "typescript/client",
		FileName: "index",
		FileExtension: "ts",
		OutputDir: filepath.Join(clientDataDir, "entities/"),
		OutputFileName: "index.ts",
		ForceWrite: true,
	},
}

func generateClientEntityOverview(data []entities.MetaData) {
	// Build the type
	overview := entities.ClientOverview{
		EntityNameTypes: "\n",
		Entities: data,
	}

	for _, entity := range data {
		overview.EntityNameTypes += " | '" + entity.TypeName + "'\n"
		overview.EntityNames += " | '" + entity.EntityName + "'\n"
		overview.EntityStoreImports = append(overview.EntityStoreImports, "import { create" + entity.TypeName + "Store } from './entities/" + entity.EntityName + "/store.gen';")
		overview.EntityStoreDefinitions = append(overview.EntityStoreDefinitions, entity.TypeName + ": create" + entity.TypeName + "Store,")
	}

	for _, templateFile := range clientOverviewTemplates {
		build(templateFile, overview)
	}
}
