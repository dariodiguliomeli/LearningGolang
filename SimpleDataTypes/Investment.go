package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func Investment() {
	capital, _ := strconv.ParseFloat(Utils.ScannerInput("Say your capital: "), 64)
	tax, _ := strconv.ParseFloat(Utils.ScannerInput("Say your tax: "), 64)
	years, _ := strconv.ParseFloat(Utils.ScannerInput("Say your years of this investment: "), 64)
	investment := (capital * tax * years) + capital
	fmt.Printf("Your gain will be: $%.2f", investment)
}
