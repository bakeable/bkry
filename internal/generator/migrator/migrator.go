package migrator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bakeable/bkry/internal/generator/types"
)


func (m *Migrator) Clean() error {
	// Compare the two directories to check which entities have been removed
	newConfigurations := m.GetNewConfigurations()
	oldConfigurations := m.GetOldConfigurations()

	// Collect entities that need to be removed
	entitiesToRemove := m.findEntitiesToRemove(newConfigurations, oldConfigurations)
	
	// Remove the old entities
	for _, entity := range entitiesToRemove {
		err := m.RemoveEntity(entity)
		if err != nil {
			return err
		}
	}


	// Save the new configurations as new generation
	err := m.SaveGeneration(newConfigurations)
	if err != nil {
		return err
	}

	return nil
}


func (m *Migrator) RemoveEntity(entity types.EntityConfig) error {
	// Remove the entity generated files
	templateFiles := types.GetSingleEntityTemplateFiles(m.TemplateFiles)
	// Remove file "database/repository/entities/{{.EntityName}}.gen.go"
	// Remove directory "models/entities/{{.EntityName}}"
	
	fmt.Println("Removing entity " + entity.EntityName, templateFiles)
	for _, tmpl := range templateFiles {
		templateFile := tmpl.DeepCopyWithVariables(map[string]string{
			"ServerDir": m.ServerDir,
			"ClientDir": m.ClientDir,
			"EntityName": entity.EntityName,
		})

		// Get the file path to delete
		filePath := filepath.Join(templateFile.OutputDir, templateFile.OutputFileName)


		// If the output dir ends with the entity name, remove the directory
		if filepath.Base(templateFile.OutputDir) == entity.EntityName {
			// Check if the directory exists
			fmt.Println("Removing directory " + templateFile.OutputDir)
			if _, err := os.Stat(templateFile.OutputDir); err == nil {
				// Remove the directory
				os.RemoveAll(templateFile.OutputDir)
			}
		} else {
			// Check if the file exists
			fmt.Println("Removing file " + filePath)
			if _, err := os.Stat(filePath); err == nil {
				// Remove the file
				os.Remove(filePath)
			}
		}
	}

	return nil
}

func (m *Migrator) SaveGeneration(newConfigurations []types.EntityConfig) error {
	// Create a new generation directory
	generationName := fmt.Sprintf("generation_%d", time.Now().Unix())
	err := os.MkdirAll(filepath.Join(m.GenerationDir, generationName), os.ModePerm)
	if err != nil {
		return err
	}

	// Save all entity configurations as .json files
	for _, entity := range newConfigurations {
		data, err := json.Marshal(entity)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(m.GenerationDir, generationName, entity.TypeName + ".json"), data, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}