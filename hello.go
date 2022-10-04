package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Gabriel"
	fmt.Println("Aaaaaaaaaaaaaaaaa", nome)
	var version = 1.1
	fmt.Println("Esse programa está na versão", version)
	fmt.Println("O version é uma variável do tipo", reflect.TypeOf(version))

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Mostrar logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	//& aponta pro endereço de uma variável
	fmt.Scan(&comando)
	fmt.Println("Vc escolheu o comando", comando)
}
