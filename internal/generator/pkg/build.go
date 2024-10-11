package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bakeable/bkry/internal/generator/types"
)

func build(templateFile types.TemplateFile, data interface{}) {
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
		"hasItems": hasItems,
		"pascalCaseToSnakeCase": pascalCaseToSnakeCase,
		"pascalCaseToCamelCase": pascalCaseToCamelCase,
	}).ParseFiles(filepath.Join(templateDir, nonGeneratedFileName) + ".tmpl"))

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
}
