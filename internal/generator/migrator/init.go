package migrator

import "github.com/bakeable/bkry/internal/generator/types"

func NewMigrator(configDir, generationDir, serverDir, clientDir string, templateFiles []types.TemplateFile) *Migrator {
	return &Migrator{
		ConfigDir:     configDir,
		GenerationDir: generationDir,
		ServerDir:     serverDir,
		ClientDir:     clientDir,
		TemplateFiles: templateFiles,
	}
}