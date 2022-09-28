package Strings

import "fmt"
import "goBases/Utils"
import "strings"

func ReformatName() {
	username := Utils.ScannerInput("Say your name: ")
	fmt.Println(strings.ToUpper(username))
	fmt.Println(strings.Title(username))
	fmt.Println(strings.ToLower(username))
}
