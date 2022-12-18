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

	for {
		exbiMenu()
		opcao := capturaOpcao()
		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
			imprimeLogs()

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
			registraLog(site, true)
		} else {
			fmt.Println("O Site: ", site, "Encontra-se inoperante no momento", resp.StatusCode)
			registraLog(site, false)
		}

	}

}

func lerSitesDoArquivo() []string {
	sites := []string{}
	arquivo, err := os.Open("sites.txt")

	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)

	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err != nil {

			if err == io.EOF {
				fmt.Println("Final de linha do arquivo")
				break

			} else {
				fmt.Println("Erro ao ler arquivo", err)
			}
		}
		sites = append(sites, linha)
		fmt.Println(linha)

	}

	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
	}
	fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "  online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivos de log", err)
	} else {
		fmt.Println(string(arquivo))
	}

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
