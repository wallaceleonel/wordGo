package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	leSitesDoArquivo()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Finalizando programa.")
			os.Exit(0)

		default:
			fmt.Println("comando executado invalido.")
			os.Exit(-1)
		}
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

		for i, ambiente := range ambientesAcademia {
			fmt.Println("testando ambiente", i, ":", ambiente)
			testaAmbiente(ambiente)
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaAmbiente(ambiente string) {
	resonse, error := http.Get(ambiente)

	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}

	if resonse.StatusCode == 200 {
		fmt.Println("Ambiente", ambiente, "carregado com sucesso!")
		registraLog(ambiente, true)

	} else {
		fmt.Println("Sit", ambiente, "está fora do ar. status code", resonse.StatusCode)
		registraLog(ambiente, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Ocorreu um erro ", error)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if error == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}

func imprimeLogs() {

	arquivo, error := ioutil.ReadFile("log.txt")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(arquivo)
}

func registraLog(ambiente string, status bool) {

	fmt.Println("Carregando os logs da aplicação.")

	arquivo, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println(error)
	}

	arquivo.WriteString((time.Now().Format("02/01/2006 15:04:05")) + "-" + ambiente + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
