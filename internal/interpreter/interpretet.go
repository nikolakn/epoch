package interpreter

import (
	"bufio"
	"epoch/internal/gps"
	"epoch/pkg/epoch"
	"fmt"
	"math"
	"os"
	"regexp"
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
		line := strings.TrimSpace(string(userData))
		if line == "q" || line == "quit" || line == "exit" {
			return
		}
		if line == "help" || line == "h" || line == "?" {
			help :=
				`
help	
	document
		s | save                save
		q | exit | quit         exit
	add/delate
		a    | add                 add new event or epoch 
		del  | delate              delate of event or epoch 
		ae   | 'add event'         add new absolute event 
		are  | 'add rel event'     add new relative event 
		aep  | 'add epoch '        add new absolute epoch 
		arep | 'add rel epoch'     add new relative epoch 
	print
		p  | print              print timeline 
		pd | 'print des'        print description of event or epoch 
		distance | dis          duration in years between start date of two event or epoch 
	edit
		r | rename | title      rename event or epoch 
		d | des                 change description of event or epoch 
		m | move                change start date of event or epoch 
		set                     set print options 
		g | gps                 geo location; position for maps
		url | u                 url of event or epoch doc 
		importance | lvl        level of importance of event or epoch 
		type                    type of event or epoch 
	`
			fmt.Println(help)
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

		if line == "gps" || line == "g" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				l1 := getfloat64("latitude")
				l2 := getfloat64("longitude")
				event.GetEpoch().GPS = gps.NewGPS(gps.Degrees(l1), gps.Degrees(l2))
			}
		}

		if line == "rename" || line == "r" || line == "title" {
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

		if line == "search title" || line == "st" {
			title := getStringInput("title")
			events := doc.SearchEventsByTitle(title)
			for _, e := range events {
				println(e.GetEpoch().Id, e.GetEpoch().Title)
			}
		}
		if line == "search des" || line == "sd" {
			title := getStringInput("description")
			events := doc.SearchEventsByDescription(title)
			for _, e := range events {
				println(e.GetEpoch().Id, e.GetEpoch().Title)
			}
		}

		if line == "add" || line == "a" {
			fmt.Print("epoch or event(ev) (defoult event) > ")
			userData, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				fmt.Println("input error: ", err)
				return
			}
			ud := string(userData)
			if ud == "event" || ud == "" || ud == "ev" {
				isRelative := yesNo("is relative start")
				if !isRelative {
					line = "ae"
				} else {
					line = "are"
				}
			} else {
				isRelative := yesNo("is relative start")
				if !isRelative {
					line = "aep"
				} else {
					line = "arep"
				}
			}
		}
		if line == "add event" || line == "ae" {
			d, m, y := getDateInput(doc.PrintOptions.YearOnly)
			h, min := 0, 0
			if doc.PrintOptions.Time {
				h, min = getTimeInput()
			}
			start := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)
			doc.AddEventWithData(start, getStringInput("title"))
		}
		if line == "add rel event" || line == "are" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				rel := getRelative()
				doc.AddRelativeEventWithData(event, rel+1, getStringInput("title"))
			}
		}
		if line == "add epoch" || line == "aep" {
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
			if end.Before(start) {
				fmt.Println("end date is before start, not added!")
			} else {
				doc.AddEpochWithData(start, end, getStringInput("title"))
			}
		}
		if line == "add rel epoch" || line == "arep" {
			event := getPArentEventByTitleorId(doc)
			if event != nil {
				rel := getRelative()
				fmt.Println("relative end:")
				rel_end := getRelative()

				doc.AddRelativeEpochWithData(event, rel, rel_end, getStringInput("title"))
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
			fmt.Print("day or date> ")
			text, _, err4 := bufio.NewReader(os.Stdin).ReadLine()
			if err4 != nil {
				fmt.Println("error: ", err4)
				continue
			}
			re := regexp.MustCompile(`\d{1,2}.\d{1,2}.(-)*\d`)
			isDate := re.Match(text)
			if isDate {
				arr := strings.Split(string(text), ".")
				d, _ = strconv.Atoi(arr[0])
				m, _ = strconv.Atoi(arr[1])
				y, _ = strconv.Atoi(arr[2])
				return d, m, y
			} else {
				if d, err4 = strconv.Atoi(string(text)); err4 != nil {
					continue
				}
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
