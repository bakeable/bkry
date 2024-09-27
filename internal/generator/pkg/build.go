package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)
type TemplateFile struct {
	TemplateDir string `json:"template_dir"`
	FileName string `json:"file_name"`
	FileExtension string `json:"file_extension"`
	OutputDir string `json:"output_dir"`
	OutputFileName string `json:"output_file_name"`
	ForceWrite bool `json:"force_write"`
	InitializeOnly bool `json:"initialize_only"`
	InputData string `json:"input_data"`
}

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

func (tf TemplateFile) DeepCopy() TemplateFile {
	return TemplateFile{
		TemplateDir: tf.TemplateDir,
		FileName: tf.FileName,
		FileExtension: tf.FileExtension,
		OutputDir: tf.OutputDir,
		OutputFileName: tf.OutputFileName,
		ForceWrite: tf.ForceWrite,
		InitializeOnly: tf.InitializeOnly,
		InputData: tf.InputData,
	}
}

func (tf TemplateFile) DeepCopyWithVariables(variables map[string]string) TemplateFile {
	replaceVariables := func (input string) string {
		for key, value := range variables {
			input = strings.ReplaceAll(input, "{" + key + "}", value)
		}
		return input
	}

	return TemplateFile{
		TemplateDir: replaceVariables(tf.TemplateDir),
		FileName: replaceVariables(tf.FileName),
		FileExtension: replaceVariables(tf.FileExtension),
		OutputDir: replaceVariables(tf.OutputDir),
		OutputFileName: replaceVariables(tf.OutputFileName),
		ForceWrite: tf.ForceWrite,
		InitializeOnly: tf.InitializeOnly,
		InputData: tf.InputData,
	}
}