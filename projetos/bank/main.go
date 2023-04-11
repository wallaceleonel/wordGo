package main

import "fmt"

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

func main() {

	contaDoJones := ContaCorrente{}
	contaDoJones.titular = "Jones Manoel"
	contaDoJones.saldoConta = 300.98

	fmt.Println(contaDoJones.Depositar(150.44))
	fmt.Println(contaDoJones.Sacar(100.98))
}
