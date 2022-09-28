package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func Division() {
	firstNumber, _ := strconv.ParseInt(Utils.ScannerInput("Say your first number: "), 10, 9)
	secondNumber, _ := strconv.ParseInt(Utils.ScannerInput("Say your second number: "), 10, 9)
	rest := firstNumber % secondNumber
	quotient := firstNumber / secondNumber
	fmt.Printf("The %d between %d give us a rest %d and a quotient %d", firstNumber, secondNumber, rest, quotient)
}
