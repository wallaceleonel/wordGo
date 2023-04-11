package contas

import "projetos/banco/clientes"

type ContaCorrente struct {
	Titular             clientes.Titular
	NumeroAgencia       int
	NumeroContaCorrente int
	SaldoConta          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {

	validaSaque := valorDoSaque > 0 && valorDoSaque <= c.SaldoConta
	if validaSaque {
		c.SaldoConta -= valorDoSaque
		return "Saque realizado com sucesso", c.SaldoConta
	} else {
		return " Saldo insuficiente", c.SaldoConta
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	validaDeposito := valorDoDeposito > 0
	if validaDeposito {
		c.SaldoConta += valorDoDeposito
		return "Deposito realizado com sucesso", c.SaldoConta
	} else {
		return "Valor de deposito invalido", c.SaldoConta
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia > 0 && valorTransferencia < c.SaldoConta {
		c.SaldoConta -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
