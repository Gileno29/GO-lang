package main

import (
	"fmt"
)

func main() {
	nome := "Gileno"
	versao := 1.1
	fmt.Println("Ola ", nome)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - sair do Programa")
	var opcao int

	//fmt.Scanf("%d", &opcao)

	//leitura de input'
	fmt.Scan(&opcao)
	fmt.Println("A operação escolhida foi", opcao)

	/*if opcao == 1 {
		fmt.Println("Monitorando...")

	} else if opcao == 2 {
		fmt.Println("Exibindo Logs")

	} else if opcao == 0 {
		fmt.Println("Saindo do programa")

	} else {
		fmt.Println("Comando não pertence as opções válidas")
	}*/

	switch opcao {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs")

	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Comando não pertence as opções válidas")
	}
}
