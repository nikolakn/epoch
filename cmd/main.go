package main

import (
	"epoch/pkg/epoch"
	"fmt"
)

func main() {

	doc := epoch.NewDocument()
	doc.LoadFromJson("test_data/test.json")
	fmt.Println(doc)
}
