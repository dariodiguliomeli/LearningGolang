package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func NewSum() {
	number, _ := strconv.ParseInt(Utils.ScannerInput("Input your integer positive number: "), 10, 9)
	result := (number * (number + 1)) / 2
	fmt.Printf("The result is: %d", result)
}
