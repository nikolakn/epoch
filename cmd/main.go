package main

import (
	"epoch/pkg/epoch"
	"fmt"
	"log"

	cmd "epoch/command"
)

func main() {
	file, print_olny, _ := cmd.Execute()
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
	if print_olny {
		doc := epoch.NewDocument(po)
		doc.LoadFromJson(file)
		fmt.Println(doc)
	}
}
