package interpreter

import (
	"bufio"
	"epoch/pkg/epoch"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

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

func getPArentEventByTitleOrId(doc *epoch.Document) epoch.Event {
	fmt.Print("parent id or title > ")
	text, _, err4 := bufio.NewReader(os.Stdin).ReadLine()
	if err4 != nil {
		fmt.Println("error: ", err4)
		return nil
	}
	if id, err := strconv.Atoi(string(text)); err == nil {
		event := doc.GetEventById(id)
		if event != nil {
			fmt.Println(doc.PrintEvent(event))
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

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}
}

func getTempPath() string {
	switch runtime.GOOS {
	case "linux":
		return "/tmp/epoch.html"
	case "windows":
		return os.TempDir() + "/epoch.html"
	case "darwin":
		return "/tmp/epoch.html"
	default:
		return "/tmp/epoch.html"
	}
}
