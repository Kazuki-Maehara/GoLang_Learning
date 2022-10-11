package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	workbook := excelize.NewFile()

	index := workbook.NewSheet("Sheet1")
	var sheetName string = workbook.GetSheetName(index)
	// fmt.Printf("%T", index)

	workbook.SetCellValue(sheetName, "A1", "Hello world.")
	for i := 1; i <= 10; i++ {
		workbook.SetCellValue("Sheet1", "B"+strconv.Itoa(i), "B"+strconv.Itoa(i))
	}

	result, err := workbook.SearchSheet(sheetName, "^B[1-9]{1}$", true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	if err := workbook.SaveAs("./Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
