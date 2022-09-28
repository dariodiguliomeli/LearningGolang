package Strings

import (
	"fmt"
	"goBases/Utils"
	"strings"
)

func PhoneNumberFormat() {
	phoneNumber := Utils.ScannerInput("Input your phone number in this format +<prefix>-<number>-<extension>: ")
	phoneNumberParts := strings.Split(phoneNumber, "-")
	fmt.Printf("The phone number is: %s", phoneNumberParts[1])
}
