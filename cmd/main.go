package main

import (
	"epoch/internal/interpreter"
	"epoch/pkg/epoch"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

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

			arr := strings.Split(runtime.Version(), ".")
			v1 := strings.TrimPrefix(arr[0], "go")
			v2, err := strconv.Atoi(arr[1])
			if err != nil {
				if v1 == "1" && v2 < 21 {
					fmt.Println("app must be compiled with go version >=go1.21.1")
				}
			}
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
