package migrator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bakeable/bkry/internal/generator/types"
)


func (m *Migrator) GetNewConfigurations() []types.EntityConfig {
	// Lees alle .json-bestanden in de configuratiemap
	files, err := os.ReadDir(m.ConfigDir)
	if err != nil {
		panic(err)
	}

	entityConfigurations := []types.EntityConfig{}
	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			entity := types.EntityConfig{}
			data, err := os.ReadFile(m.ConfigDir + file.Name())
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(data, &entity)
			if err != nil {
				panic(err)
			}
			entityConfigurations = append(entityConfigurations, entity)
		}
	}

	return entityConfigurations
}

func (m *Migrator) GetOldConfigurations() []types.EntityConfig {
	// Read directories in the generation directory
	dirs, err := os.ReadDir(m.GenerationDir)
	if err != nil {
		fmt.Println("Error reading generations directory")
		return m.GetNewConfigurations()
	}

	// Get the latest generation directory by name
	var latestDir os.DirEntry
	for _, dir := range dirs {
		if dir.IsDir() {
			if latestDir == nil {
				latestDir = dir
			} else if dir.Name() > latestDir.Name() {
				latestDir = dir
			}
		}
	}

	// Read all .json files in the latest generation directory
	files, err := os.ReadDir(filepath.Join(m.GenerationDir, latestDir.Name()) + "/")
	if err != nil {
		return m.GetNewConfigurations()
	}

	// Collect all entity configurations
	entityConfigurations := []types.EntityConfig{}
	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			entity := types.EntityConfig{}
			data, err := os.ReadFile(filepath.Join(m.GenerationDir, latestDir.Name(), file.Name()))
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(data, &entity)
			if err != nil {
				panic(err)
			}
			entityConfigurations = append(entityConfigurations, entity)
		}
	}

	return entityConfigurations
}

func (m *Migrator) findEntitiesToAdd(newConfigurations []types.EntityConfig, oldConfigurations []types.EntityConfig) []types.EntityConfig {
	// Collect new entities or entities that have been removed
	entitiesToAdd := []types.EntityConfig{}

	// Compare the two slices
	for _, newEntity := range newConfigurations {
		found := false
		for _, oldEntity := range oldConfigurations {
			if newEntity.TypeName == oldEntity.TypeName {
				found = true
				break
			}
		}

		if !found {
			entitiesToAdd = append(entitiesToAdd, newEntity)
		}
	}

	return entitiesToAdd
}

func (m *Migrator) findEntitiesToRemove(newConfigurations []types.EntityConfig, oldConfigurations []types.EntityConfig) []types.EntityConfig {
	// Collect new entities or entities that have been removed
	entitiesToRemove := []types.EntityConfig{}

	// Compare the two slices
	for _, oldEntity := range oldConfigurations {
		found := false
		for _, newEntity := range newConfigurations {
			if newEntity.TypeName == oldEntity.TypeName {
				found = true
				break
			}
		}

		if !found {
			entitiesToRemove = append(entitiesToRemove, oldEntity)

			// If it is localised, remove the localisation file too
			if oldEntity.IsLocalizable {
				entitiesToRemove = append(entitiesToRemove, types.EntityConfig{
					TypeName: oldEntity.TypeName + "Localisation",
					EntityName: oldEntity.EntityName + "_localisation",
				})
			}
		}
	}

	return entitiesToRemove
}

func getEntity(typeName string, entities []types.EntityConfig) *types.EntityConfig {
	for _, entity := range entities {
		if entity.TypeName == typeName {
			return &entity
		}
	}
	return nil
}

func getField(fieldName string, fields []types.FieldConfig) *types.FieldConfig {
	for _, field := range fields {
		if field.FieldName == fieldName {
			return &field
		}
	}
	return nil
}

func (m *Migrator) findEntitiesToModify(newConfigurations []types.EntityConfig, oldConfigurations []types.EntityConfig) []types.EntityConfig {
	// Collect new entities or entities that have been removed
	entitiesToModify := []types.EntityConfig{}

	// Compare the two slices
	for _, newEntity := range newConfigurations {
		modified := false
		oldEntity := getEntity(newEntity.TypeName, oldConfigurations)
		if oldEntity != nil {
			for _, newField := range newEntity.Fields {
				oldField := getField(newField.FieldName, oldEntity.Fields)
				if oldField != nil {
					if newField.IsModified(*oldField) {
						modified = true
						break
					}
				}
			}
		}

		if modified {
			entitiesToModify = append(entitiesToModify, newEntity)
		}
	}

	return entitiesToModify
}