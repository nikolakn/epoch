package epoch

import (
	jd "epoch/internal/julian"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (doc *Document) ExportJson(file string) {
	doc.Savejson(file)
}

func fileNameStrip(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}

func (doc *Document) ExportHtml(file string) {
	var body string
	//zoomStr := strconv.Itoa(doc.PrintOptions.Zoom)
	body += `<div class="container">
	<h2>Epoch '` + fileNameStrip(doc.FileName) + `</h2>
	<button  onclick="plus()" type="button">+</button>
	<button  onclick="minus()"" type="button">-</button>

	<ul class="responsive-table">
		<li class="table-header">
		<div class="col col-1">Start</div>
		<div class="col col-2">End</div>
		<div class="col col-3">Duration</div>
		<div class="col col-4">Name</div>
		<div class="col col-5">Description</div>
		<div class="col col-6">Url</div>
		<div class="col col-7">Map</div>
		</li>

	`
	var pre Event
	for _, e := range doc.Events {
		start := jd.JDToTime(e.GetStart())
		end := jd.JDToTime(e.GetDuration())
		days := math.Abs(e.GetDuration() - e.GetStart())
		years := days / JDYear
		months := 12.0 * years
		d2 := fmt.Sprintf("%.2f years", years)
		if months < 1 {
			d2 = fmt.Sprintf("%0.2f days", days)
		} else if years < 1 {
			d2 = fmt.Sprintf("%0.2f months", months)
		}

		if e.Relative() || e.EndRelative() {
			duration := e.GetDuration() / JDYear
			iy := int(duration)
			y_ostatak := 12.0 * float64(duration-float64(iy))
			d2 = fmt.Sprintf("%0.2f years ~(%d years and %0.2f months)", duration, iy, y_ostatak)
			end = jd.JDToTime(e.GetStart() + e.GetDuration())
		}

		if pre != nil && doc.PrintOptions.Zoom > 0 {
			e2 := pre.GetStart()
			e1 := e.GetStart()
			diff := math.Abs(e2 - e1)
			raz := int(diff / float64(doc.PrintOptions.Zoom))
			years := diff / JDYear
			months := 12.0 * years
			dd := fmt.Sprintf("%.1f years", years)
			if months < 1 {
				dd = fmt.Sprintf("%0.1f days", diff)
			} else if years < 1 {
				dd = fmt.Sprintf("%0.1f months", months)
			}
			if raz*50 > 20 {
				body += "\t\t<div  class='separator' style='align-items: center;justify-content: center;display: flex; height:" + strconv.Itoa(raz*50) + "px'>" + dd + "</div>\n"
			}
		}
		pre = e
		body += "\t <li class='table-row'>\n"
		body += "\t\t<div class='col col-1' data-label='Start'>" + fmt.Sprintf(" %d.%d.%d", start.Day(), start.Month(), start.Year()) + "</div>\n"
		if e.GetDuration() == 0 {
			body += "\t\t<div class='col col-2' data-label='End'></div>\n"
		} else {
			body += "\t\t<div class='col col-2' data-label='End'>" + fmt.Sprintf(" %d.%d.%d", end.Day(), end.Month(), end.Year()) + "</div>\n"
		}
		if e.GetDuration() == 0 {
			body += "\t\t<div class='col col-3' data-label='Duration'></div>\n"

		} else {
			body += "\t\t<div class='col col-3' data-label='Duration'>" + d2 + "</div>\n"
		}
		body += "\t\t<div class='col col-4' data-label='Name'>" + fmt.Sprintf("%s", e.GetEpoch().Title) + "</div>\n"
		body += "\t\t<div class='col col-5' data-label='Description'>" + fmt.Sprintf("%s", e.GetEpoch().Description) + "</div>\n"
		if e.GetEpoch().Url != "" {
			body += "\t\t<div class='col col-6' data-label='Url'><a target='_blank' rel='noopener' href='" + fmt.Sprintf("%s", e.GetEpoch().Url) + "'>link</a></div>\n"
		} else {
			body += "\t\t<div class='col col-6' data-label='Url'></div>\n"
		}
		if e.GetEpoch().GPS.Latitude != 0 {
			body += "\t\t<div class='col col-7' data-label='Map'><a target='_blank' rel='noopener' href='https://www.osmap.uk/#10/" + e.GetEpoch().GPS.PrintForMAp() + "'>map</a></div>\n"
		} else {
			body += "\t\t<div class='col col-7' data-label='Map'></div>\n"
		}
		body += "\t</li>\n"

	}

	body += "\t</ul>\n"
	body += "\t</div>\n"
	js := `

		function plus() {
			var elements = document.getElementsByClassName('separator');
			for (const element of elements) {
				var h = element.clientHeight
				console.log(h)
				element.style.height = (h*2)+"px";
			}
		}

		function minus() {
			var elements = document.getElementsByClassName('separator');
			for (const element of elements) {
				var h = element.clientHeight
				if (h<20) {
					return
				}
			}
			for (const element of elements) {
				var h = element.clientHeight
				console.log(h)
				element.style.height = (h/2)+"px";
			}
		}
	`
	html := `
  <!DOCTYPE html>
  <html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <meta http-equiv="X-UA-Compatible" content="ie=edge">
	  <title>Epoch</title>
	  <script>` + js + `</script>
	  <style>
	  body {
		font-family: "lato", sans-serif;
	  }
	  
	  .container {
		max-width: 90%;
		margin-left: auto;
		margin-right: auto;
		padding-left: 10px;
		padding-right: 10px;
	  }
	  
	  h2 {
		font-size: 26px;
		margin: 20px 0;
		text-align: center;
	  }
	  h2 small {
		font-size: 0.5em;
	  }
	  
	  .responsive-table li {
		border-radius: 4px;
		padding: 16px 16px;
		display: flex;
		justify-content: space-between;
		margin-bottom: 16px;
		font-size: 12px;
	  }
	  .responsive-table .table-header {
		background-color: #95a5a6;
		font-size: 14px;
		text-transform: uppercase;
		letter-spacing: 0.03em;
	  }
	  .responsive-table .table-row {
		background-color: #ffffff;
		box-shadow: 0px 0px 9px 0px rgba(0, 0, 0, 0.1);
	  }
	  .responsive-table .col-1 {
		flex-basis: 10%;
	  }
	  .responsive-table .col-2 {
		flex-basis: 10%;
	  }
	  .responsive-table .col-3 {
		flex-basis: 15%;
	  }
	  .responsive-table .col-4 {
		flex-basis: 20%;
	  }
	  .responsive-table .col-5 {
		flex-basis:35%;
	  }
	  .responsive-table .col-6 {
		flex-basis: 5%;
	  }
	  .responsive-table .col-7 {
		flex-basis: 5%;
	  }
	  @media all and (max-width: 767px) {
		.responsive-table .table-header {
		  display: none;
		}
		.responsive-table li {
		  display: block;
		}
		.responsive-table .col {
		  flex-basis: 100%;
		}
		.responsive-table .col {
		  display: flex;
		  padding: 5px 0;
		}
		.responsive-table .col:before {
		  color: #6c7a89;
		  padding-right: 10px;
		  content: attr(data-label);
		  flex-basis: 20%;
		  text-align: left;
		}
	  }
	  </style>
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
