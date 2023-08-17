package interpreter

import (
	"bufio"
	"epoch/internal/gps"
	"epoch/pkg/epoch"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
		if line == "q" || line == "quit" || line == "exit" {
			return
		}
		if line == "help" || line == "h" || line == "?" {
			fmt.Println("p; print\t-\tprint timeline")
			fmt.Println("a; add\t\t-\tadd new event or epoch")
			fmt.Println("r; rename\t-\trename event or epoch")
			fmt.Println("d; des\t\t-\tchange description of event or epoch")
			fmt.Println("m; move\t\t-\tchange start date of event or epoch")
			fmt.Println("set\t\t-\tset print options")
			fmt.Println("pd; print des\t-\tprint description of event or epoch")
			fmt.Println("distance; dis\t-\tduration in years between start date of two event or epoch")
			fmt.Println("location; gps\t-\tgeo location of  event or epoch")
			fmt.Println("del; delate\t-\tdelate of event or epoch")

			fmt.Println("url; u\t-\turl of event or epoch doc")
			fmt.Println("importance; lvl\t-\tlevel of importance of event or epoch")
			fmt.Println("type; \t-\ttype of event or epoch")

			fmt.Println("q; exit; quit\t-\texit")
			continue
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
			doc.PrintOptions.Description = yesNo("display description")

		}
		if line == "des" || line == "d" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				des := getStringInput("description")
				event.GetEpoch().Description = des
			}
		}

		if line == "url" || line == "u" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				t := getStringInput("url")
				event.GetEpoch().Url = t
			}
		}
		if line == "importance" || line == "lvl" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				t := getInt("Importance")
				event.GetEpoch().Importance = t
			}
		}
		if line == "type" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				t := getInt("Tyepe")
				event.GetEpoch().Type = t
			}
		}

		if line == "gps" || line == "location" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				l1 := getfloat64("latitude")
				l2 := getfloat64("longitude")
				event.GetEpoch().GPS = gps.NewGPS(gps.Degrees(l1), gps.Degrees(l2))
			}
		}

		if line == "des" || line == "d" {
			//#TODO curently text is readed until new line fix to read until ctrl-D
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				des := getStringInput("description")
				event.GetEpoch().Description = des
			}
		}

		if line == "rename" || line == "r" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				t := getStringInput("title")
				event.GetEpoch().Title = t
			}
		}
		if line == "move" || line == "m" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				if event.GetEpoch().IsRelative {
					rel := getRelative()
					doc.MoveStartRel(event, rel)
					if event.GetDuration() != 0 {
						event.SetEnd(rel + event.GetDuration())

					}
				} else {
					d, m, y := getDateInput(doc.PrintOptions.YearOnly)
					h, min := 0, 0
					if doc.PrintOptions.Time {
						h, min = getTimeInput()
					}
					start := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
					doc.MoveStartAps(event, start)
					if event.GetDuration() != 0 {
						event.SetEnd(event.GetDuration())

					}
				}
			}
		}

		if line == "print des" || line == "pd" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				fmt.Println(event.GetEpoch().Description)
			}
		}

		if line == "delate" || line == "del" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				doc.DeleteEvent(event)
			}
		}

		if line == "distance" || line == "dis" {
			event1 := getPArentEventByTitleorId(doc)
			event2 := getPArentEventByTitleorId(doc)
			if event1 != nil && event2 != nil {
				d := math.Abs(event2.GetStart() - event1.GetStart())
				years := float64(d) / epoch.JDYear
				iy := int(years)
				y_ostatak := 12.0 * float64(years-float64(iy))
				fmt.Printf("diffrence in years %f years ~(%d years and %0.2f months)\n", years, iy, y_ostatak)
			}
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
					d, m, y := getDateInput(doc.PrintOptions.YearOnly)
					h, min := 0, 0
					if doc.PrintOptions.Time {
						h, min = getTimeInput()
					}
					start := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
					doc.AddEventWithData(start, getStringInput("title"))
				} else {
					event := getPArentEventByTitleorId(doc)
					if event != nil {
						rel := getRelative()
						doc.AddRelativeEventWithData(event, rel+1, getStringInput("title"))
					}

				}
			} else {
				isRelative := yesNo("is relative start")
				if !isRelative {
					d, m, y := getDateInput(doc.PrintOptions.YearOnly)
					h, min := 0, 0
					if doc.PrintOptions.Time {
						h, min = getTimeInput()
					}
					start := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)

					fmt.Println("enter end date:")
					h, min = 0, 0
					d, m, y = getDateInput(doc.PrintOptions.YearOnly)
					if doc.PrintOptions.Time {
						h, min = getTimeInput()
					}
					end := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
					doc.AddEpochWithData(start, end, getStringInput("title"))
				} else {
					event := getPArentEventByTitleorId(doc)
					if event != nil {
						fmt.Println("relative start from parent event (title): ", event.GetEpoch().Title)
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
	s := strings.TrimSpace(string(text))
	return string(s)
}

func getfloat64(msg string) float64 {
	var v float64
	fmt.Print(msg + " > ")
	_, err := fmt.Scanf("%f\n", &v)
	if err != nil {
		fmt.Println("error: ", err)
		return 0
	}
	return v
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

func getInt(msg string) int {
	var id int
	fmt.Print(msg + " > ")
	_, err := fmt.Scanf("%d\n", &id)
	if err != nil {
		fmt.Println("error: ", err)
		return 0
	}
	return id
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

func getPArentEventByTitleorId(doc *epoch.Document) epoch.Event {
	fmt.Print("parent id or title > ")
	text, _, err4 := bufio.NewReader(os.Stdin).ReadLine()
	if err4 != nil {
		fmt.Println("error: ", err4)
		return nil
	}
	if id, err := strconv.Atoi(string(text)); err == nil {
		event := doc.GetEventbuId(id)
		if event != nil {
			fmt.Println("Selected:", event.GetEpoch().Title)
		}
		return event
	}
	event := doc.GetEventbyTitle(string(text))
	if event != nil {
		fmt.Println("Selected:", event.GetEpoch().Title)
	}
	return event
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

func getDateInput(yearsOnly bool) (int, int, int) {
	for true {
		var d, m, y int
		if !yearsOnly {
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
		} else {
			d = 1
			m = 1
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

func getTimeInput() (int, int) {
	for true {
		var h, m int

		fmt.Print("hour > ")
		_, err := fmt.Scanf("%2d\n", &h)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		fmt.Print("minute > ")
		_, err2 := fmt.Scanf("%2d\n", &m)
		if err2 != nil {
			fmt.Println("error: ", err2)
			continue
		}

		return h, m
	}
	return 0, 0
}
