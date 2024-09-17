package generator

import (
	"github.com/bakeable/bkry/internal/generator/entities"
)

func GenerateRoutes(entities []entities.MetaData) {
	for _, entity := range entities {
		// Generate standard route file
		generateRoutes(entity)
	}

	// Generate api file
	generateApi(entities)
}

func generateRoutes(entity entities.MetaData) {
	// Generate default route file
	build(TemplateFile{
		TemplateDir: "golang/routes",
		FileName: "default",
		FileExtension: "go",
		OutputDir: ServerDir + "api/" + entity.PackagePath + "/routes/",
	}, entity)

	// Generate custom route file
	build(TemplateFile{
		TemplateDir: "golang/routes",
		FileName: "custom",
		FileExtension: "go",
		OutputDir: ServerDir + "api/" + entity.PackagePath + "/routes/",
	}, entity)
}


func generateApi(entities []entities.MetaData) {
	build(TemplateFile{
		TemplateDir: "golang",
		FileName: "api",
		FileExtension: "go",
		OutputDir: ServerDir + "api/",
		ForceWrite: true,
	}, entities)
}