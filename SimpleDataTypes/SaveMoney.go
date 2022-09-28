package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func SaveMoney() {
	const INTEREST = 0.04
	capital, _ := strconv.ParseFloat(Utils.ScannerInput("Say your capital: "), 64)
	firstYearSave := capital + (capital * INTEREST)
	secondYearSave := firstYearSave + (firstYearSave * INTEREST)
	thirdYearSave := secondYearSave + (secondYearSave * INTEREST)
	fmt.Printf("Your gain after first year will be: $%.2f \n", firstYearSave)
	fmt.Printf("Your gain after second year will be: $%.2f \n", secondYearSave)
	fmt.Printf("Your gain after third year will be: $%.2f \n", thirdYearSave)
}
