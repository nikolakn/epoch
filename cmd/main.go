package main

import (
	"epoch/internal/interpreter"
	"epoch/pkg/epoch"
	"fmt"
	"path/filepath"

	cmd "epoch/command"
)

func main() {
	file, outFile, print_olny, _ := cmd.Execute()
	if file == "" {
		fmt.Println("invalid file name, making tmp file tmp_epoch.json!")
		file = "tmp_epoch.json"
	}

	po := epoch.PrintOptions{
		Flags:    true,
		YearOnly: false,
		Time:     false,
		Duration: true,
		GPS:      true,
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
		if ext == ".xlsx" {
			doc.SaveExcel(outFile)
		}

		return
	}
	if print_olny {
		//print document on console and exit
		doc := epoch.NewDocument(po, file)
		doc.LoadFromJson(file)
		fmt.Println(doc)
	} else {
		//interpreter
		interpreter.NewInterpreter(file)
	}
}
