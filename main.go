package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exbiMenu()
		opcao := capturaOpcao()
		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")

		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não pertence as opções válidas")
			os.Exit(-1)
		}

	}
}

func exibeIntroducao() {
	nome := "Gileno"
	versao := 1.1
	fmt.Println("Ola ", nome)
	fmt.Println("Este programa está na versão: ", versao)

}

func capturaOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	fmt.Println("A operação escolhida foi", opcao)

	return opcao

}

func exbiMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://www.linkedin.com.br/"
	sites[1] = "https://youtube.com.br/"
	sites[2] = "https://facebook.com.br/"

	resp, _ := http.Get(sites)
	if resp.StatusCode == 200 {
		fmt.Println("O siste:", site, "encontra-se operante no momento")
	} else {
		fmt.Println("O Site: ", site, "Encontra-se inoperante no momento", resp.StatusCode)
	}
}

//fmt.Scanf("%d", &opcao)

//leitura de input'

/*if opcao == 1 {
	fmt.Println("Monitorando...")

} else if opcao == 2 {
	fmt.Println("Exibindo Logs")

} else if opcao == 0 {
	fmt.Println("Saindo do programa")

} else {
	fmt.Println("Comando não pertence as opções válidas")
}*/
