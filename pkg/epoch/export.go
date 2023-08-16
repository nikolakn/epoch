package epoch

import (
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
	  <th>Company</th>
	  <th>Contact</th>
	  <th>Country</th>
	</tr>
	<tr>
    <td>Alfreds Futterkiste</td>
    <td>Maria Anders</td>
    <td>Germany</td>
	</tr>
	<tr>
		<td>Centro comercial Moctezuma</td>
		<td>Francisco Chang</td>
		<td>Mexico</td>
	</tr>
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
