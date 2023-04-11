package main

import "fmt"

type ContaCorrente struct {
	titular             string
	numeroAgencia       int
	numeroContaCorrente int
	saldoConta          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	validaSaque := valorDoSaque > 0 && valorDoSaque <= c.saldoConta
	if validaSaque {
		c.saldoConta -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return " Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	validaDeposito := valorDoDeposito > 0
	if validaDeposito {
		c.saldoConta += valorDoDeposito
		return "Deposito realizado com sucesso"
	} else {
		return "Valor de deposito invalido"
	}
}

func main() {

	contaDoJones := ContaCorrente{}
	contaDoJones.titular = "Jones Manoel"
	contaDoJones.saldoConta = 300.98

	fmt.Println(contaDoJones.Depositar(150.44))
	fmt.Println(contaDoJones.Sacar(100.98))
}
