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

	contaDoGo := contaCorrente{
		"Go",
		1313,
		1231344,
		212.98}

	var contaDoJones *contaCorrente
	contaDoJones = new(contaCorrente)
	contaDoJones.titular = "Jones Manoel"
	contaDoJones.numeroAgencia = 1313
	contaDoJones.numeroContaCorrente = 21211212
	contaDoJones.saldoConta = 131313.13

	fmt.Println(*contaDoJones)
	fmt.Println(contaDoGoberto)
	fmt.Println(contaDoGo)
}
