package main

import (
	"fmt"
)

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

func printTemperatureConverter() {
	drawTable("°C", "°F", ctof)
	fmt.Println()
	drawTable("°F", "°C", ftoc)
}

func drawTable(header1, header2 string, tableTd tableLine) {
	initialTemp := -40
	finalTemp := 100
	fmt.Println(line)
	fmt.Printf(rowFormat, header1, header2)
	fmt.Println(line)
	for i := initialTemp; i <= finalTemp; i = i + 5 {
		cell1, cell2 := tableTd(i)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
	fmt.Println()
}

type tableLine func(tableTd int) (string, string)
type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func ctof(value int) (string, string) {
	var c celsius = celsius(value)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func ftoc(value int) (string, string) {
	var f fahrenheit = fahrenheit(value)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}
