package SimpleDataTypes

import "fmt"
import "goBases/Utils"

func HelloUser() {
	username := Utils.ScannerInput("Say your name: ")
	fmt.Printf("¡Hello %s!\n", username)
}
