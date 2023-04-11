package main

import (
	"C:\.git\wordGo\projetos\bank\account"
	"fmt"
)

func main() {

	contaDoJones := account.ContaCorrente{}
	contaDoJones.Titular = "Jones Manoel"
	contaDoJones.SaldoConta = 600.98

	contaDoManoel := account.ContaCorrente{}
	contaDoManoel.Titular = "Manoel Jones"
	contaDoManoel.SaldoConta = 300.98

	fmt.Println(contaDoJones.Depositar(150.44))
	fmt.Println(contaDoJones.Sacar(100.98))

	fmt.Println(contaDoManoel.Depositar(150.44))
	fmt.Println(contaDoManoel.Sacar(100.98))

	status := contaDoJones.Transferir(50.55, &contaDoManoel)

	fmt.Println(status)
	fmt.Print(contaDoJones.SaldoConta)
	fmt.Print(contaDoManoel.SaldoConta)

}
