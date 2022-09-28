package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func IMC() {
	weight, _ := strconv.ParseFloat(Utils.ScannerInput("Say your weight: "), 64)
	height, _ := strconv.ParseFloat(Utils.ScannerInput("Say your height: "), 64)
	imc := weight / (height * height)
	fmt.Printf("Your IMC is: %.2f", imc)
}
