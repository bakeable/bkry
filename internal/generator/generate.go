package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	entities "github.com/bakeable/bkry/internal/generator/entities"
	generator "github.com/bakeable/bkry/internal/generator/scripts"
)

//go:generate go run generate.go

func main() {
	// Example data for an entity
	data := entities.GetMetaData()

	// Preprocess all entities
	localizedEntities := []entities.MetaData{}
	for i, entity := range data {
		fmt.Println("\n\nGenerating code for entity: " + entity.TypeName)
		
		// If entity is localizable, generate a child instance of the entity
		if entity.IsLocalizable {
			localizedEntity := generator.Preprocess(entity.DeepCopy(), true)
			localizedEntities = append(localizedEntities, localizedEntity)
		}

		// Preprocess all entities
		data[i] = generator.Preprocess(entity.DeepCopy(), false)
	}

	// Add localized entities to the data
	data = append(data, localizedEntities...)

	// Generate the client side scripts
	generator.GenerateClient(data)

	// Generate the back end scripts
	generator.GenerateBackend(data)
	
	// Generate the entity config
	generator.GenerateConfig(data)

	// Generate the entity repository
	generator.GenerateRepository(data)

	// Generate the entity actions
	generator.GenerateOperations(data)

	// Generate controllers
	generator.GenerateControllers(data)

	// Generate routes
	generator.GenerateRoutes(data)

	// Change directory to ClientDir
	runFormatter(generator.ClientDir)
	runFormatter(generator.BackEndDir)
}


func runFormatter(dir string) {
	// Get current directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		return
	}

	// Create absolute path to dir
	dir = filepath.Join(cwd, dir)

	// Change directory to dir
	err = os.Chdir(dir)
	if err != nil {
		fmt.Println("Failed to change directory:", err)
		return
	}

	// Run "npm run format" in dir
	exec.Command("npm", "run", "format").Run()

	// Change directory back to original directory
	err = os.Chdir(cwd)
}