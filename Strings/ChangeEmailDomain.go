package Strings

import (
	"fmt"
	"goBases/Utils"
	"strings"
)

func ChangeEmailDomain() {
	email := Utils.ScannerInput("Say your email: ")
	emailParts := strings.Split(email, "@")
	fmt.Printf("%s@ceu.es", emailParts[0])
}
