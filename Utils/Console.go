package Utils

import (
	"bufio"
	"fmt"
	"os"
)

func ScannerInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	return input
}
