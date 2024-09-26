package generator

import (
	"github.com/bakeable/bkry/input/entities"
)


type RepositoryInterfaceData struct {
	Entities []entities.MetaData
}
func GenerateRepository(entitySlice []entities.MetaData) {
	// Generate index file for the entity
	build(TemplateFile{
		TemplateDir: "golang/repository",
		FileName: "repo_interface",
		FileExtension: "go",
		OutputDir: ServerDir + "data/repository/entities/",
		OutputFileName: "index.go",
		ForceWrite: true,
	},  RepositoryInterfaceData{ Entities: entitySlice })

	// Generate repository files for the entity
	for _, entity := range entitySlice {
		build(TemplateFile{
			TemplateDir: "golang/repository",
			FileName: "repository",
			FileExtension: "go",
			OutputDir: ServerDir + "data/repository/entities/",
			OutputFileName: entity.EntityName + ".gen.go",
			ForceWrite: true,
		}, entity)
	}
}