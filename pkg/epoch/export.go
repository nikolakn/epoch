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
	body += `<div class="container">
	<h2>Epoch</h2>
	<ul class="responsive-table">
		<li class="table-header">
		<div class="col col-1">Start</div>
		<div class="col col-2">End</div>
		<div class="col col-3">Duration</div>
		<div class="col col-4">Name</div>
		<div class="col col-5">Description</div>
		<div class="col col-6">Url</div>
		</li>

	`
	for _, e := range doc.Events {
		start := jd.JDToTime(e.GetStart())
		end := jd.JDToTime(e.GetDuration())
		duration := (e.GetDuration() - e.GetStart()) / JDYear
		iy := int(duration)
		y_ostatak := 12.0 * float64(duration-float64(iy))
		d2 := fmt.Sprintf("%.1f y ~(%d y %0.1f m)", duration, iy, y_ostatak)

		if e.Relative() || e.EndRelative() {
			duration = e.GetDuration() / JDYear
			iy := int(duration)
			y_ostatak := 12.0 * float64(duration-float64(iy))
			d2 = fmt.Sprintf("%0.2f years ~(%d years and %0.2f months)", duration, iy, y_ostatak)
			end = jd.JDToTime(e.GetStart() + e.GetDuration())
		}

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
		body += "\t\t<div class='col col-6' data-label='Url'>" + fmt.Sprintf("%s", e.GetEpoch().Url) + "</div>\n"

		body += "\t</li>\n"

	}

	body += "\t</ul>\n"
	body += "\t</div>\n"
	html := `
  <!DOCTYPE html>
  <html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <meta http-equiv="X-UA-Compatible" content="ie=edge">
	  <title>Epoch</title>
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
		flex-basis: 20%;
	  }
	  .responsive-table .col-4 {
		flex-basis: 20%;
	  }
	  .responsive-table .col-5 {
		flex-basis:20%;
	  }
	  .responsive-table .col-6 {
		flex-basis: 20%;
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
