package migrator

import "github.com/bakeable/bkry/internal/generator/types"

type Migrator struct {
	ConfigDir              string
	GenerationDir          string
	ServerDir 			   string
	ClientDir 			   string
	TemplateFiles		   []types.TemplateFile
}