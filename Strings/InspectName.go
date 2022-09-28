package Strings

import (
	"fmt"
	"goBases/Utils"
	"strings"
)

func InspectName() {
	username := Utils.ScannerInput("Say your name: ")
	username = strings.ToUpper(username)
	fmt.Printf("The name %s has %d letters", username, len(username))
}
