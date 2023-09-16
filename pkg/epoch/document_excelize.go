package epoch

import (
	jd "epoch/internal/julian"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (doc *Document) SaveExcel(name string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.SetCellValue("Sheet1", "A1", "Id")
	f.SetCellValue("Sheet1", "B1", "Flags")
	f.SetCellValue("Sheet1", "C1", "Start")
	f.SetCellValue("Sheet1", "D1", "End")
	f.SetCellValue("Sheet1", "E1", "Parent")
	f.SetCellValue("Sheet1", "F1", "Title")
	f.SetColWidth("Sheet1", "F", "F", 30)
	f.SetCellValue("Sheet1", "G1", "Duration (days)")
	f.SetCellValue("Sheet1", "H1", "GPS")
	f.SetColWidth("Sheet1", "H", "H", 15)
	f.SetCellValue("Sheet1", "I1", "Description")
	f.SetColWidth("Sheet1", "I", "I", 50)
	for index, event := range doc.Events {
		doc.PrintRow(event, f, index)
	}
	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
}

func (doc Document) PrintRow(e Event, f *excelize.File, row int) string {
	index := 1
	text := ""
	coord := getRawLetter(index) + fmt.Sprint(row+2)
	f.SetCellValue("Sheet1", coord, e.GetEpoch().Id)
	coord = getRawLetter(index+1) + fmt.Sprint(row+2)
	f.SetCellValue("Sheet1", coord, doc.PrintFlags(e))
	coord = getRawLetter(index+2) + fmt.Sprint(row+2)
	f.SetCellValue("Sheet1", coord, doc.PrintStartClear(e))
	if e.GetDuration() != 0 {
		if e.Relative() || e.EndRelative() {
			time := jd.JDToTime(e.GetStart() + e.GetDuration())
			coord := getRawLetter(index+3) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, doc.PrintEndClear(time))
			coord = getRawLetter(index+4) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, e.GetEpoch().Id)
			coord = getRawLetter(index+5) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, doc.PrintTitle(e))
			coord = getRawLetter(index+6) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, e.GetDuration())
		} else {
			time := jd.JDToTime(e.GetDuration())
			coord := getRawLetter(index+3) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, doc.PrintEndClear(time))
			coord = getRawLetter(index+4) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, e.GetEpoch().Id)
			coord = getRawLetter(index+5) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, doc.PrintTitle(e))
			coord = getRawLetter(index+6) + fmt.Sprint(row+2)
			f.SetCellValue("Sheet1", coord, e.GetDuration()-e.GetStart())
		}
	} else {
		coord = getRawLetter(index+4) + fmt.Sprint(row+2)
		f.SetCellValue("Sheet1", coord, e.GetEpoch().Id)
		coord = getRawLetter(index+5) + fmt.Sprint(row+2)
		f.SetCellValue("Sheet1", coord, doc.PrintTitle(e))
	}
	coord = getRawLetter(index+7) + fmt.Sprint(row+2)
	f.SetCellValue("Sheet1", coord, doc.PrintGPSclear(e))
	coord = getRawLetter(index+8) + fmt.Sprint(row+2)
	f.SetCellValue("Sheet1", coord, doc.PrintDescriptionClear(e))
	return text
}

func getRawLetter(row int) (result string) {
	arr := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return fmt.Sprintf("%s", arr[row-1])
}
