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
}
