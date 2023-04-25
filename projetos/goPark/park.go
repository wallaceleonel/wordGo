package main

import (
	"fmt"
	"time"
)

type carro struct {
	placa   string
	horario time.Time
}

var estacionamento = make(map[string]carro)

const taxaHora = 14.0

func entrada(placa string) {
	if _, ok := estacionamento[placa]; ok {
		fmt.Println("Este carro já está estacionado.")
	} else {
		estacionamento[placa] = carro{placa: placa, horario: time.Now()}
		fmt.Printf("O carro de placa %s entrou no estacionamento.\n", placa)
	}
}

func saida(placa string) {
	if carro, ok := estacionamento[placa]; ok {
		tempoEstacionado := time.Since(carro.horario)
		taxa := taxaHora * float64(tempoEstacionado.Hours())
		delete(estacionamento, placa)
		fmt.Printf("O carro de placa %s saiu do estacionamento. Taxa a pagar: R$%.2f\n", placa, taxa)
	} else {
		fmt.Println("Este carro não está estacionado no momento.")
	}
}

func main() {
	entrada("ABC1234")
	time.Sleep(1 * time.Hour)
	saida("ABC1234")
	entrada("DEF5678")
	time.Sleep(2 * time.Hour)
	saida("ABC1234")
	saida("DEF5678")
}
