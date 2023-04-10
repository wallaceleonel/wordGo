package main

import "fmt"

func main() {
	nome := "Gonewson"
	idade := 9
	versao := 1.1

	fmt.Println("ola sr.", nome, "cujo idade é ", idade)
	fmt.Println("Esse programa está na versão.", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi ", comando)

	//if comando == 1 {
	//	fmt.Println("iniciando monitoramento.")
	//} else if comando == 2 {
	//	fmt.Println("Carregando os logs da aplicação.")
	//} else if comando == 0 {
	//	fmt.Println("Finalizando programa.")
	//} else {
	//	fmt.Println("comando executado invalido.")
	//}

	switch comando {

	case 1:
		fmt.Println("iniciando monitoramento.")
	case 2:
		fmt.Println("Carregando os logs da aplicação.")
	case 0:
		fmt.Println("Finalizando programa.")

	default:
		fmt.Println("comando executado invalido.")
	}
}
