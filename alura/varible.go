package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {

	case 1:
		iniciarMonitoramento()
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
	fmt.Println("")
	return comandoLido
}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("iniciando monitoramento.")
	var ambientesAcademia = []string{"https://playmove-ead-producao.azurewebsites.net/", "https://playmove-ead-stage.azurewebsites.net/"}

	for i := 0; i < monitoramentos; i++ {

		for i, site := range ambientesAcademia {
			fmt.Println("testando ambiente", i, ":", site)
			testaAmbiente(site)
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaAmbiente(ambientes string) {
	resonse, _ := http.Get(ambientes)

	if resonse.StatusCode == 200 {
		fmt.Println("Ambiente", ambientes, "carregado com sucesso!")
	} else {
		fmt.Println("Sit", ambientes, "está fora do ar. status code", resonse.StatusCode)
	}
}
