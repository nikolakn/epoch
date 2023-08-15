package interpreter

import (
	"bufio"
	"epoch/pkg/epoch"
	"fmt"
	"os"
	"time"
)

func NewInterpreter(fileName string) {
	po := epoch.PrintOptions{
		Flags:    true,
		YearOnly: false,
		Time:     false,
		Duration: true,
		GPS:      false,
		Id:       true,
	}

	doc := epoch.NewDocument(po, fileName)
	if fileName != "" {
		doc.LoadFromJson(fileName)
	}
	for true {
		fmt.Print("Epoch > ")
		userData, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			fmt.Println("input error: ", err)
			return
		}
		line := string(userData)
		if line == "q" || line == "quit" || line == "close" || line == "exit" {
			return
		}
		if line == "save" || line == "s" {
			doc.Savejson(fileName)
		}
		if line == "print" || line == "p" {
			fmt.Println(doc)
		}

		if line == "add" || line == "a" {
			fmt.Print("epoch or event (empty for event) > ")
			userData, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				fmt.Println("input error: ", err)
				return
			}
			line := string(userData)
			if line == "event" || line == "" {
				d, m, y := getDateInput()
				start := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
				doc.AddEventWithData(start, getStringInput("title"))
			} else {

			}
		}
	}
}
func getStringInput(msg string) string {
	fmt.Print(msg + " > ")
	text, _, err4 := bufio.NewReader(os.Stdin).ReadLine()
	if err4 != nil {
		fmt.Println("error: ", err4)
		return ""
	}
	return string(text)
}

func getDateInput() (int, int, int) {
	var d, m, y int
	fmt.Print("day > ")
	_, err := fmt.Scanf("%2d\n", &d)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Print("month > ")
	_, err2 := fmt.Scanf("%2d\n", &m)
	if err2 != nil {
		fmt.Println("error: ", err2)
	}
	fmt.Print("year > ")
	_, err3 := fmt.Scanf("%d\n", &y)
	if err3 != nil {
		fmt.Println("error: ", err3)
	}
	return d, m, y
}
