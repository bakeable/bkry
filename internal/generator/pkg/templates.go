package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bakeable/bkry/internal/generator/types"
)

// LoadConfig reads the config file from a template directory
func LoadTemplatesConfig(configPath string) ([]types.TemplateConfigItem, error) {
	var config []types.TemplateConfigItem
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// RetrieveTemplateFiles handles globbing for output_dir and direct file path for output_file
func RetrieveTemplateFiles(rootDir string, item types.TemplateConfigItem) ([]string, error) {
	var templateFiles []string
	// Check if it's a glob pattern (output_dir defined)
	if item.OutputDir != "" {
		matches, err := filepath.Glob(filepath.Join(rootDir, item.Path))
		if err != nil {
			return nil, err
		}
		templateFiles = append(templateFiles, matches...)
	} else if item.OutputFile != "" {
		// Direct file, check if it exists
		fullPath := filepath.Join(rootDir, item.Path)
		if _, err := os.Stat(fullPath); err == nil {
			templateFiles = append(templateFiles, fullPath)
		} else {
			return nil, err
		}
	}
	return templateFiles, nil
}

// ProcessTemplates traverses directories and reads the configuration for code generation
func ProcessTemplates(rootDir string) ([]types.TemplateFile, error) {
	// Create a slice of types.TemplateFile configurations
	var templateFiles []types.TemplateFile = make([]types.TemplateFile, 0)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if we have a config file in the current directory
		if info.IsDir() {
			configPath := filepath.Join(path, "bkry.templates.config.json")
			if _, err := os.Stat(configPath); err == nil {
				configItems, err := LoadTemplatesConfig(configPath)
				if err != nil {
					return err
				}

				// Process each config item
				for _, item := range configItems {
					fmt.Printf("Processing template config: %s\n", item.Path)

					// Retrieve the relevant template files
					templateFilePaths, err := RetrieveTemplateFiles(path, item)
					if err != nil {
						return err
					}

					// Print the list of template files found
					for _, filePath := range templateFilePaths {
						fmt.Printf("Found template file: %s\n", filePath)
						// Get the directory of the template file
						templateDir := filepath.Dir(filePath)

						// Get the template file name
						templateFileName := strings.Split(filepath.Base(filePath), ".tmpl")[0]

						// Get the extension of the template file
						templateExtension := strings.Join(strings.Split(templateFileName, ".")[1:], ".")

						// Remove the extension from the template file name
						templateFileName = strings.Split(templateFileName, ".")[0]

						templateFiles = append(templateFiles, types.TemplateFile{
							TemplateDir: 	templateDir,
							FileName:   	templateFileName,
							FileExtension: 	templateExtension,
							OutputDir:  	item.OutputDir,
							OutputFileName: item.OutputFile,
							ForceWrite: 	item.ForceWrite,
							InitializeOnly: item.InitializeOnly,
							InputData: 		item.Data,
						})
					}

					// Additional processing logic for templates would go here
				}
			}
		}

		return nil
	})

	return templateFiles, err
}
