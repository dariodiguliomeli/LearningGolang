package Strings

import (
	"fmt"
	"goBases/Utils"
	"strings"
)

func UpperSingleCharacter() {
	text := Utils.ScannerInput("Say some text: ")
	vocal := Utils.ScannerInput("Say your favorite vocal in this text: ")
	text = strings.ReplaceAll(text, vocal, strings.ToUpper(vocal))
	fmt.Print(text)
}
