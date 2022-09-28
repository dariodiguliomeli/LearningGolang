package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func CalculateSalary() {
	hours, _ := strconv.ParseFloat(Utils.ScannerInput("Say your hours worked: "), 64)
	price, _ := strconv.ParseFloat(Utils.ScannerInput("Say your price per hours worked: "), 64)
	salary := hours * price
	fmt.Printf("Your salary is $%.2f", salary)
}
