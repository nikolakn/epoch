package epoch

import (
	jd "epoch/internal/julian"
	"fmt"
	"log"
	"os"
)

func (doc *Document) ExportJson(file string) {
	doc.Savejson(file)
}

func (doc *Document) ExportHtml(file string) {
	var body string
	body += `
	<table>
	<tr>
	  <th>Start</th>
	  <th>End</th>
	  <th>Duration</th>
	  <th>Name</th>
	  <th>Description</th>
	  <th>Url</th>
	</tr>
	`
	for _, e := range doc.Events {
		start := jd.JDToTime(e.GetStart())
		end := jd.JDToTime(e.GetDuration())
		duration := (e.GetDuration() - e.GetStart()) / JDYear
		iy := int(duration)
		y_ostatak := 12.0 * float64(duration-float64(iy))
		d2 := fmt.Sprintf("%.2f years ~(%d years and %0.2f months)\n", duration, iy, y_ostatak)

		if e.Relative() || e.EndRelative() {
			duration = e.GetDuration() / JDYear
			iy := int(duration)
			y_ostatak := 12.0 * float64(duration-float64(iy))
			d2 = fmt.Sprintf("%f years ~(%d years and %0.2f months)\n", duration, iy, y_ostatak)
			end = jd.JDToTime(e.GetStart() + e.GetDuration())
		}

		body += "<tr>\n"
		body += "<td>" + fmt.Sprintf(" %d.%d.%d", start.Day(), start.Month(), start.Year()) + "</td>\n"
		if e.GetDuration() == 0 {
			body += "<td></td>\n"
		} else {
			body += "<td>" + fmt.Sprintf(" %d.%d.%d", end.Day(), end.Month(), end.Year()) + "</td>\n"
		}
		if e.GetDuration() == 0 {
			body += "<td></td>\n"

		} else {
			body += "<td>" + d2 + "</td>\n"
		}
		body += "<td>" + fmt.Sprintf("%s", e.GetEpoch().Title) + "</td>\n"
		body += "<td>" + fmt.Sprintf("%s", e.GetEpoch().Description) + "</td>\n"
		body += "<td>" + fmt.Sprintf("%s", e.GetEpoch().Url) + "</td>\n"

		body += "</tr>\n"

	}

	body += `
	</table>
	`
	html := `
  <!DOCTYPE html>
  <html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <meta http-equiv="X-UA-Compatible" content="ie=edge">
	  <title>Epoch</title>
	</head>
	<body>

	` + body + `
	</body>
  </html>
  `
	err_write := os.WriteFile(file, []byte(html), 0644)
	if err_write != nil {
		log.Println("error write to file:  ", err_write)
		return
	}
}
