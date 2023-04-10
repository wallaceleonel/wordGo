package main

import "fmt"

type contaCorrente struct {
	titular             string
	numeroAgencia       int
	numeroContaCorrente int
	saldoConta          float64
}

func main() {

	contaDoGoberto := contaCorrente{
		titular:             "Goberto Fagundes",
		numeroAgencia:       1313,
		numeroContaCorrente: 12344566,
		saldoConta:          12231.98}
	fmt.Println(contaDoGoberto)

	contaDoGo := contaCorrente{
		"Go", 1313, 1231344, 212.98}
	fmt.Println(contaDoGo)
}
