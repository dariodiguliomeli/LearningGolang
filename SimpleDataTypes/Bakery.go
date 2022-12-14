package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func Bakery() {
	const BREADPRICE = 3.49
	const OLDBREADDISCOUNT = 0.6
	oldBreadSold, _ := strconv.ParseFloat(Utils.ScannerInput("Say your bread sold amount: "), 64)
	fmt.Printf("The usual price of bread is $%.2f\n", BREADPRICE)
	fmt.Printf("The old bread discount is %.2f%%\n", OLDBREADDISCOUNT*100)
	partialPrice := oldBreadSold * BREADPRICE
	totalPrice := partialPrice - (partialPrice * OLDBREADDISCOUNT)
	fmt.Printf("The total price sold today is $%.2f\n", totalPrice)
}
