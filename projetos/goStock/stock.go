package main

import (
	"fmt"
)

type produto struct {
	quantidade int
	preco      float64
}

var estoque = make(map[string]produto)

func adicionar_produto(nome string, quantidade int, preco float64) {
	if _, ok := estoque[nome]; ok {
		estoque[nome].quantidade += quantidade
	} else {
		estoque[nome] = produto{quantidade: quantidade, preco: preco}
	}
}

func remover_produto(nome string, quantidade int) {
	if _, ok := estoque[nome]; ok {
		if estoque[nome].quantidade >= quantidade {
			estoque[nome].quantidade -= quantidade
		} else {
			fmt.Println("Não há quantidade suficiente deste produto no estoque.")
		}
	} else {
		fmt.Println("Este produto não está cadastrado no estoque.")
	}
}

func exibir_estoque() {
	fmt.Println("Estoque:")
	for nome, info := range estoque {
		fmt.Printf("%s - Quantidade: %d - Preço: R$%.2f\n", nome, info.quantidade, info.preco)
	}
}

// Teste do controle de estoque
func main() {
	adicionar_produto("Caneta", 10, 2.50)
	adicionar_produto("Lápis", 5, 1.50)
	exibir_estoque()

	remover_produto("Caneta", 3)
	exibir_estoque()

	remover_produto("Caneta", 10)

	adicionar_produto("Caneta", 5, 2.75)
	exibir_estoque()
}
