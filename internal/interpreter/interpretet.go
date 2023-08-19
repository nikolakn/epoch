package interpreter

import (
	"bufio"
	"epoch/internal/gps"
	jd "epoch/internal/julian"
	"epoch/pkg/epoch"
	"fmt"
	"math"
	"os"
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

		printRange(line, doc)
		flags(line, doc)
		move(line, doc)

		if line == "q" || line == "quit" || line == "exit" {
			return
		}
		if line == "help" || line == "h" || line == "?" {
			fmt.Println(HELP)
			continue
		}
		if line == "save" || line == "s" {
			doc.Savejson(fileName)
		}
		if line == "print" || line == "p" {
			fmt.Println(doc)
		}

		if line == "des" || line == "d" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				des := getStringInput("description")
				event.GetEpoch().Description = des
			}
		}

		if line == "url" || line == "u" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				t := getStringInput("url")
				event.GetEpoch().Url = t
			}
		}

		if line == "open map" {
			event := getPArentEventByTitleOrId(doc)
			if event == nil {
				fmt.Println("event does not exist")
			} else {
				if event.GetEpoch().GPS.Latitude == 0 {
					fmt.Println("location data for event missing, use gps command to set location data for map")
				} else {
					openBrowser("https://www.osmap.uk/#10/" + event.GetEpoch().GPS.PrintForMAp())
				}
			}
		}

		if line == "open url" {
			event := getPArentEventByTitleOrId(doc)
			if event == nil {
				fmt.Println("event does not exist")
			} else {
				if event.GetEpoch().Url == "" {
					fmt.Println("Url for event missing, use url command to set link")
				} else {
					openBrowser(event.GetEpoch().Url)
				}
			}
		}

		if line == "open html" {
			doc.ExportHtml(getTempPath())
			openBrowser("file:" + getTempPath())
		}

		if line == "importance" || line == "lvl" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				t := getInt("Importance")
				event.GetEpoch().Importance = t
			}
		}
		if line == "type" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				t := getInt("Type")
				event.GetEpoch().Type = t
			}
		}

		if line == "gps" || line == "g" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				l1 := getfloat64("latitude")
				l2 := getfloat64("longitude")
				event.GetEpoch().GPS = gps.NewGPS(gps.Degrees(l1), gps.Degrees(l2))
			}
		}

		if line == "rename" || line == "r" || line == "title" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				t := getStringInput("title")
				event.GetEpoch().Title = t
			}
		}

		if line == "print des" || line == "pd" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				fmt.Println(event.GetEpoch().Description)
			}
		}

		if line == "delate" || line == "del" {
			event := getPArentEventByTitleOrId(doc)
			if event != nil {
				doc.DeleteEvent(event)
			}
		}

		if line == "distance" || line == "dis" {
			event1 := getPArentEventByTitleOrId(doc)
			event2 := getPArentEventByTitleOrId(doc)
			if event1 != nil && event2 != nil {
				d := math.Abs(event2.GetStart() - event1.GetStart())
				years := float64(d) / epoch.JDYear
				months := 12.0 * years

				d2 := fmt.Sprintf("%.2f years", years)
				if months < 1 {
					d2 = fmt.Sprintf("%0.2f days", d)
				} else if years < 1 {
					d2 = fmt.Sprintf("%0.2f months", months)
				}
				fmt.Printf("diffrence %s\n", d2)

			}
		}

		if line == "search title" || line == "st" {
			title := getStringInput("title")
			events := doc.SearchEventsByTitle(title)
			for _, e := range events {
				fmt.Println(doc.PrintEvent(e))
			}
		}
		if line == "search des" || line == "sd" {
			title := getStringInput("description")
			events := doc.SearchEventsByDescription(title)
			for _, e := range events {
				fmt.Println(doc.PrintEvent(e))
			}
		}

		add(line, doc)
	}
}

func flags(line string, doc *epoch.Document) {
	if line == "set" {
		doc.PrintOptions.Flags = yesNo("display flags")
		doc.PrintOptions.Id = yesNo("display id")
		doc.PrintOptions.Time = yesNo("display time")
		doc.PrintOptions.GPS = yesNo("display gps")
		doc.PrintOptions.Duration = yesNo("display duration")
		doc.PrintOptions.YearOnly = yesNo("display year only")
		doc.PrintOptions.Description = yesNo("display description")
	}

	if line == "show flags" {
		doc.PrintOptions.Flags = true
	}
	if line == "show id" {
		doc.PrintOptions.Id = true
	}
	if line == "show time" {
		doc.PrintOptions.Time = true
	}
	if line == "show gps" {
		doc.PrintOptions.GPS = true
	}
	if line == "show duration" {
		doc.PrintOptions.Duration = true
	}
	if line == "show year only" {
		doc.PrintOptions.YearOnly = true
	}
	if line == "show description" {
		doc.PrintOptions.Description = true
	}

	if line == "hide flags" {
		doc.PrintOptions.Flags = false
	}
	if line == "hide id" {
		doc.PrintOptions.Id = false
	}
	if line == "hide time" {
		doc.PrintOptions.Time = false
	}
	if line == "hide gps" {
		doc.PrintOptions.GPS = false
	}
	if line == "hide duration" {
		doc.PrintOptions.Duration = false
	}
	if line == "hide year only" {
		doc.PrintOptions.YearOnly = false
	}
	if line == "hide description" {
		doc.PrintOptions.Description = false
	}

	if line == "set zoom" {
		zoom := getInt("zoom (one line represent distance in days for example zoom=30 one line is 30 days)")
		doc.PrintOptions.Zoom = zoom
	}
}

func printRange(line string, doc *epoch.Document) {
	if line == "print range" || line == "pr" {
		d, m, y := getDateInput(doc.PrintOptions.YearOnly)
		d_end, m_end, y_end := getDateInput(doc.PrintOptions.YearOnly)
		start := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		end := time.Date(y_end, time.Month(m_end), d_end, 0, 0, 0, 0, time.UTC)
		s := jd.TimeToJD(start)
		e := jd.TimeToJD(end)
		fmt.Println(s, e)
		var pre epoch.Event
		for _, event := range doc.Events {
			date := event.GetStart()

			if date >= s && date < e {
				if pre != nil && doc.PrintOptions.Zoom > 0 {
					e2 := pre.GetStart()
					e1 := event.GetStart()
					diff := math.Abs(e2 - e1)
					raz := int(diff / float64(doc.PrintOptions.Zoom))
					fmt.Print(strings.Repeat("|\n", raz))
				}
				pre = event
				fmt.Println(doc.PrintEvent(event))
			}
		}
	}
}

func add(line string, doc *epoch.Document) {
	if line == "add" || line == "a" {
		fmt.Print("epoch or event(ev) (default event) > ")
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
		event := getPArentEventByTitleOrId(doc)
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
		event := getPArentEventByTitleOrId(doc)
		if event != nil {
			rel := getRelative()
			fmt.Println("relative end:")
			rel_end := getRelative()

			doc.AddRelativeEpochWithData(event, rel, rel_end, getStringInput("title"))
		}
	}
}

func move(line string, doc *epoch.Document) {
	if line == "move" || line == "m" {
		event := getPArentEventByTitleOrId(doc)
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
}
