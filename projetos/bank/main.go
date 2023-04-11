package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular             string
	numeroAgencia       int
	numeroContaCorrente int
	saldoConta          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {

	validaSaque := valorDoSaque > 0 && valorDoSaque <= c.saldoConta
	if validaSaque {
		c.saldoConta -= valorDoSaque
		return "Saque realizado com sucesso", c.saldoConta
	} else {
		return " Saldo insuficiente", c.saldoConta
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	validaDeposito := valorDoDeposito > 0
	if validaDeposito {
		c.saldoConta += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldoConta
	} else {
		return "Valor de deposito invalido", c.saldoConta
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia > 0 && valorTransferencia < c.saldoConta {
		c.saldoConta -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {

	contaDoJones := ContaCorrente{}
	contaDoJones.titular = "Jones Manoel"
	contaDoJones.saldoConta = 600.98

	contaDoManoel := ContaCorrente{}
	contaDoManoel.titular = "Manoel Jones"
	contaDoManoel.saldoConta = 300.98

	fmt.Println(contaDoJones.Depositar(150.44))
	fmt.Println(contaDoJones.Sacar(100.98))

	fmt.Println(contaDoManoel.Depositar(150.44))
	fmt.Println(contaDoManoel.Sacar(100.98))

	status := contaDoJones.Transferir(50.55, &contaDoManoel)

	fmt.Println(status)
	fmt.Print(contaDoJones.saldoConta)
	fmt.Print(contaDoManoel.saldoConta)

}
