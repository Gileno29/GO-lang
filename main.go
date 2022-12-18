package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3

const delay = 5

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
	sites := lerSitesDoArquivo()

	fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		for i := 0; i < len(sites); i++ {
			testaSite(sites[i])

		}
		time.Sleep(delay * time.Second)

	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Erro ao moniorar o site: %s", site+"\n")

	} else {
		if resp.StatusCode == 200 {
			fmt.Println("O siste:", site, "encontra-se operante no momento")
		} else {
			fmt.Println("O Site: ", site, "Encontra-se inoperante no momento", resp.StatusCode)
		}

	}

}

func lerSitesDoArquivo() []string {
	sites := []string{}
	arquivo, err := os.Open("sites.txt")

	//arquivo, err := ioutil.ReadFile("sites.txt")

	fmt.Println(arquivo)

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)

	}
	leitor := bufio.NewReader(arquivo)
	linha, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Println("ocorreu um erro ao ler o arquivo com o bufferio")
	}

	fmt.Println(linha)
	return sites

}

/*
func exibirNomes() {
	nomes := []string{"Dougglas", "Daniel", "Bernado"}
	fmt.Println(nomes)
}

//fmt.Scanf("%d", &opcao)

//leitura de input'

/*if opcao == 1 {1

	fmt.Println("Monitorando...")

} else if opcao == 2 {
	fmt.Println("Exibindo Logs")

} else if opcao == 0 {
	fmt.Println("Saindo do programa")

} else {
	fmt.Println("Comando não pertence as opções válidas")
}*/
