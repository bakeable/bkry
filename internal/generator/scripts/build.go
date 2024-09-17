package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bakeable/bkry/internal/generator/entities"
)
type TemplateFile struct {
	TemplateDir string `json:"template_dir"`
	FileName string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	OutputDir string `json:"output_dir"`
	OutputFileName string `json:"output_file_name"`
	ForceWrite bool `json:"force_write"`
	InitializeOnly bool `json:"initialize_only"`
	InputData InputData `json:"input_data"`
}

type InputData string
const (
	Entity InputData = "entity"
	Entities InputData = "entities"
)

func build(templateFile TemplateFile, data interface{}) {
	// Get the template data
	templateDir := templateFile.TemplateDir
	fileName := templateFile.FileName
	fileExtension := templateFile.FileExtension
	outputDir := templateFile.OutputDir

	// Non-generated name
	nonGeneratedFileName := fileName + "." + fileExtension
	generatedFileName := fileName + ".gen." + fileExtension
	if templateFile.InitializeOnly {
		generatedFileName = nonGeneratedFileName
	}

	// Parse the template from file
	tmpl := template.Must(template.New(nonGeneratedFileName + ".tmpl").Funcs(template.FuncMap{
		"notIn": notIn,
		"alreadyGenerated": alreadyGenerated,
	}).ParseFiles(filepath.Join("templates", templateDir, nonGeneratedFileName) + ".tmpl"))

	// If an output file name is specified, use that
	if templateFile.OutputFileName != "" {
		nonGeneratedFileName = templateFile.OutputFileName
		generatedFileName = templateFile.OutputFileName
	}

	// Create config files for every entity
	filePath := filepath.Join(outputDir, nonGeneratedFileName)

	// Create necessary directories if necessary
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Check if a standard file already exists (without .gen)
	if _, err := os.Stat(filePath); err == nil && !templateFile.ForceWrite {
		fmt.Println(filePath + " exists and has been customized. Skipping.")
		return
	} else {
		filePath = filepath.Join(outputDir, generatedFileName)
	}

	// Execute the template with the entity data
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated " + filePath)

	// Check if the data is a slice or not
	if _, ok := data.([]entities.MetaData); ok {
		templateFile.InputData = Entities
	} else {
		templateFile.InputData = Entity
	}

	addToJSON(TemplateDir + "template_files.json", templateFile)
}

// readAndUpdateJSON reads the JSON file, adds the new object, and returns the updated array
func addToJSON(filePath string, newTemplateFile TemplateFile) {
	var objects []TemplateFile

	// Try to read the file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		// If the file doesn't exist, initialize an empty array
		if os.IsNotExist(err) {
			objects = []TemplateFile{}
		} else {
			panic(err)
		}
	} else {
		// If the file exists, parse the JSON into the array
		err = json.Unmarshal(file, &objects)
		if err != nil {
			panic(err)
		}
	}

	// Add the new object to the array
	objects = append(objects, newTemplateFile)

	// Save the JSON file
	err = saveJSON(filePath, objects)
	if err != nil {
		panic(err)
	}
}

// saveJSON saves the updated array of objects to the JSON file
func saveJSON(filePath string, objects []TemplateFile) error {
	data, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("File saved successfully")
	return nil
}