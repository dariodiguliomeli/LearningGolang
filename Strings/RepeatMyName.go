package Strings

import (
	"fmt"
	"goBases/Utils"
	"strconv"
	"strings"
)

func RepeatMyName() {
	username := Utils.ScannerInput("Say your name: ")
	amount, _ := strconv.ParseInt(Utils.ScannerInput("Say your amount to repeat: "), 10, 64)
	repeated := strings.Repeat(fmt.Sprintf("%s\n", username), int(amount))
	fmt.Println(repeated)
}
