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
		if line == "set" {
			doc.PrintOptions.Flags = yesNo("display flags")
			doc.PrintOptions.Id = yesNo("display id")
			doc.PrintOptions.Time = yesNo("display time")
			doc.PrintOptions.GPS = yesNo("display gps")
			doc.PrintOptions.Duration = yesNo("display duration")
			doc.PrintOptions.YearOnly = yesNo("display year only")

		}
		if line == "add" || line == "a" {
			fmt.Print("epoch (ep) or event(ev) (empty for event) > ")
			userData, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				fmt.Println("input error: ", err)
				return
			}
			line := string(userData)
			if line == "event" || line == "" || line == "ev" {
				isRelative := yesNo("is relative start")
				if !isRelative {
					d, m, y := getDateInput()
					start := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
					doc.AddEventWithData(start, getStringInput("title"))
				} else {
					id := getParentId()
					event, _ := doc.GetEventbuId(id)
					if event != nil {
						rel := getRelative()
						doc.AddRelativeEventWithData(event, rel+1, getStringInput("title"))
					}

				}
			} else {
				isRelative := yesNo("is relative start")
				if !isRelative {
					d, m, y := getDateInput()
					start := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
					fmt.Println("enter end date:")
					d, m, y = getDateInput()
					end := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)
					doc.AddEpochWithData(start, end, getStringInput("title"))
				} else {
					id := getParentId()
					event, _ := doc.GetEventbuId(id)
					if event != nil {
						fmt.Println("relative start from parent event (id,title): ", id, event.GetEpoch().Title)
						rel := getRelative()
						fmt.Println("relative end:")
						rel_end := getRelative()

						doc.AddRelativeEpochWithData(event, rel, rel_end, getStringInput("title"))
					}
				}
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

func yesNo(msg string) bool {
	fmt.Print(msg + " y/(n) > ")
	text, _, err4 := bufio.NewReader(os.Stdin).ReadLine()
	if err4 != nil {
		fmt.Println("error: ", err4)
		return false
	}
	if string(text) == "yes" || string(text) == "y" {
		return true
	} else {
		return false
	}
}

func getParentId() int {
	var id int
	fmt.Print("parent id > ")
	_, err := fmt.Scanf("%d\n", &id)
	if err != nil {
		fmt.Println("error: ", err)
		return 0
	}
	return id
}

func getRelative() float64 {
	var y float64
	fmt.Print("relative in years > ")
	_, err := fmt.Scanf("%f\n", &y)
	if err != nil {
		fmt.Println("error: ", err)
		return 0
	}
	return y * epoch.JDYear
}

func getDateInput() (int, int, int) {
	for true {
		var d, m, y int
		fmt.Print("day > ")
		_, err := fmt.Scanf("%2d\n", &d)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		fmt.Print("month > ")
		_, err2 := fmt.Scanf("%2d\n", &m)
		if err2 != nil {
			fmt.Println("error: ", err2)
			continue
		}
		fmt.Print("year > ")
		_, err3 := fmt.Scanf("%d\n", &y)
		if err3 != nil {
			fmt.Println("error: ", err3)
			continue
		}
		return d, m, y
	}
	return 0, 0, 0
}
