package Strings

import (
	"fmt"
	"goBases/Utils"
)

func RevertText() {
	text := Utils.ScannerInput("Say some text to see it reverted: ")
	var result string
	for _, character := range text {
		result = string(character) + result
	}
	fmt.Print(result)
}
