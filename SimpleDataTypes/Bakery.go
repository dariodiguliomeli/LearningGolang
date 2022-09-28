package SimpleDataTypes

import (
	"fmt"
	"goBases/Utils"
	"strconv"
)

func Bakery() {
	// Una panadería vende barras de pan a $3.49 cada una. El pan que no es el día tiene un descuento del 60%.
	// Escribir un programa que comience leyendo el número de barras vendidas que no son del día.
	// Después el programa debe mostrar el precio habitual de una barra de pan,
	// el descuento que se le hace por no ser fresca y el coste final total.
	const BREADPRICE = 3.49
	const OLDBREADDISCOUNT = 0.6
	oldBreadSold, _ := strconv.ParseFloat(Utils.ScannerInput("Say your bread sold amount: "), 64)
	fmt.Printf("The usual price of bread is $%.2f\n", BREADPRICE)
	fmt.Printf("The old bread discount is %.2f\n", OLDBREADDISCOUNT*100)
	partialPrice := oldBreadSold * BREADPRICE
	totalPrice := partialPrice - (partialPrice * OLDBREADDISCOUNT)
	fmt.Printf("The total price sold today is $%.2f\n", totalPrice)

}
