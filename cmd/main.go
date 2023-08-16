package main

import (
	"epoch/internal/interpreter"
	"epoch/pkg/epoch"
	"fmt"
	"log"
	"path/filepath"

	cmd "epoch/command"
)

func main() {
	file, outFile, print_olny, _ := cmd.Execute()
	if file == "" {
		log.Println("invalid file name")
		return
	}

	po := epoch.PrintOptions{
		Flags:    true,
		YearOnly: false,
		Time:     false,
		Duration: true,
		GPS:      false,
		Id:       true,
	}

	if outFile != "" {
		ext := filepath.Ext(outFile)
		doc := epoch.NewDocument(po, file)
		doc.LoadFromJson(file)

		if ext == ".json" {
			doc.ExportJson(outFile)
		}

		if ext == ".html" {
			doc.ExportHtml(outFile)
		}
		return
	}
	if print_olny {
		doc := epoch.NewDocument(po, file)
		doc.LoadFromJson(file)
		fmt.Println(doc)
	} else {
		//interpreter
		interpreter.NewInterpreter(file)
	}
}
