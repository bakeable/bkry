package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bakeable/bkry/internal/generator/migrator"
	"github.com/bakeable/bkry/internal/generator/preprocessor"
	"github.com/bakeable/bkry/internal/generator/types"
)

var InputDir = "/Users/robin/Github/bkry/input/"
var GoAppName = "github.com/bakeable/bkry"

func LoadConfig(configPath string) (types.Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return types.Config{}, err
	}

	var config types.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	// Adjust paths
	config.TemplateDir = filepath.Join(InputDir, config.TemplateDir)
	config.ServerDir = filepath.Join(config.OutputDir, config.ServerDir)
	config.ClientDir = filepath.Join(config.OutputDir, config.ClientDir)
	config.BackEndDir = filepath.Join(config.OutputDir, config.BackEndDir)

	return config, nil
}

func Run() {
	// Read the bkry.config.json file in the input directory
	configPath := InputDir + "bkry.config.json"
	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error reading config file", configPath)
		panic(err)
	}


	// Process templates
	templateFiles, err := ProcessTemplates(config.TemplateDir)
	if err != nil {
		fmt.Println(err)
	}

	// Get migrator
	m := migrator.NewMigrator(
		InputDir + "config/",
		InputDir + "generations/",
		config.ServerDir,
		config.ClientDir,
		templateFiles,
	)
	m.Clean()

	// Get preprocessor 
	p := preprocessor.NewPreprocessor("/Users/robin/Github/bkry/input/config/")
	overview := p.GetOverview()
	
	for _, tmpl := range templateFiles {
		templateFile := tmpl.DeepCopyWithVariables(map[string]string{
			"ServerDir": config.ServerDir,
			"ClientDir": config.ClientDir,
			"BackEndDir": config.BackEndDir,
		})
		
		
		if templateFile.InputData == "entity_overview" {
			// Build the template
			build(templateFile, overview)
		} else if templateFile.InputData == "all_entities" {
			// Build the template
			build(templateFile, overview.Entities)
		} else {
			for _, entity := range overview.Entities {
				// Replace entity-specific variables
				entityTemplateFile := templateFile.DeepCopyWithVariables(map[string]string{
					"EntityName": entity.EntityName,
				})
				
				// Build the template
				build(entityTemplateFile, entity)
			}
		}
	}

	// Walk through all generated files and replace GoAppName by config.GoAppName
	err = filepath.Walk(config.OutputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Replace GoAppName
			file = []byte(strings.ReplaceAll(string(file), GoAppName, config.GoAppName))

			// Write the file
			err = os.WriteFile(path, file, os.ModePerm)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}