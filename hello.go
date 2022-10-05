package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	falaOi()
	mostraMenu()

	//if comando == 1 {
	//	fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//	fmt.Println("Exibindo logs...")
	//} else if comando == 0 {
	//	fmt.Println("Tô saindo...")
	//} else {
	//	fmt.Println("Comando inválido!")
	//}

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Tô saindo...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido!")
		os.Exit(-1)
	}
}

func falaOi() {
	nome := "Gabriel"
	fmt.Println("Aaaaaaaaaaaaaaaaa", nome)
	var version = 1.1
	fmt.Println("Esse programa está na versão", version)
	fmt.Println("O version é uma variável do tipo", reflect.TypeOf(version))
}

func mostraMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Mostrar logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	//& aponta pro endereço de uma variável
	fmt.Scan(&comandoLido)
	fmt.Println("Vc escolheu o comandoLido", comandoLido)

	return comandoLido
}
