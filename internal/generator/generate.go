package main

import generator "github.com/bakeable/bkry/internal/generator/pkg"

//go:generate go run generate.go

func main() {
	generator.Run()
}