package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var OutputDir = "/Users/robin/Github/bkry/output/"
var ServerDir = "/Users/robin/Github/bkry/internal/server/"
var TemplateDir = "/Users/robin/Github/bkry/input/templates/"

type ConfigItem struct {
	Path       string `json:"path"`
	OutputDir  string `json:"output_dir,omitempty"`
	OutputFile string `json:"output_file,omitempty"`
	Data       string `json:"data"`
	InitializeOnly bool `json:"initialize_only,omitempty"`
	ForceWrite bool `json:"force_write,omitempty"`
}

// LoadConfig reads the config file from a template directory
func LoadConfig(configPath string) ([]ConfigItem, error) {
	var config []ConfigItem
	file, err := ioutil.ReadFile(configPath)
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
func RetrieveTemplateFiles(rootDir string, item ConfigItem) ([]string, error) {
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
func ProcessTemplates(rootDir string) ([]TemplateFile, error) {
	// Create a slice of TemplateFile configurations
	var templateFiles []TemplateFile = make([]TemplateFile, 0)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if we have a config file in the current directory
		if info.IsDir() {
			configPath := filepath.Join(path, "bkry.config.json")
			if _, err := os.Stat(configPath); err == nil {
				configItems, err := LoadConfig(configPath)
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

						templateFiles = append(templateFiles, TemplateFile{
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

func Run() {
	// Process templates
	templateFiles, err := ProcessTemplates(TemplateDir)
	if err != nil {
		fmt.Println(err)
	}

	// Get overview
	overview := GetOverview()
	
	for _, tmpl := range templateFiles {
		templateFile := tmpl.DeepCopyWithVariables(map[string]string{
			"ServerDir": ServerDir,
			"ClientDir": ClientDir,
			"BackEndDir": BackEndDir,
		})
		
		
		fmt.Println(templateFile)
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
				fmt.Println(entityTemplateFile)
				
				// Build the template
				build(entityTemplateFile, entity)
			}
		}
	}
}