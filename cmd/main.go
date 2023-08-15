package main

import (
	"epoch/pkg/epoch"
	"fmt"
)

func main() {
	po := epoch.PrintOptions{
		Flags:    true,
		YearOnly: false,
		Time:     false,
		Duration: true,
		GPS:      false,
		Id:       true,
	}
	doc := epoch.NewDocument(po)
	doc.LoadFromJson("test_data/test_ww2.json")
	fmt.Println(doc)

	doc.LoadFromJson("test_data/test.json")

	fmt.Println(doc)
}
