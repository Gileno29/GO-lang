package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Gileno"
	var versao float32 = 1.1
	var idade int //caso não seja declarado valor para as variaveis no GO o valor inicial delas é 0 caso seja inteiro, caso seja uma string vai ser um texto vazio
	fmt.Println("olá, sr", nome, "sua iadade é: ", idade)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("O tipo da variavel é", reflect.TypeOf(nome))
}

///OUTRA FORMA DE SE DECLARAR VARIAVEIS
//nome: = "seu nome"
