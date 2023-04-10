package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("iniciando monitoramento.")
	case 2:
		fmt.Println("Carregando os logs da aplicação.")
	case 0:
		fmt.Println("Finalizando programa.")
		os.Exit(0)

	default:
		fmt.Println("comando executado invalido.")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Gonewson"
	idade := 9
	versao := 1.1

	fmt.Println("ola sr.", nome, "cujo idade é ", idade)
	fmt.Println("Esse programa está na versão.", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}
