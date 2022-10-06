package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const tempoEntreMonitoramentos = 5

func main() {
	//if comando == 1 {
	//	fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//	fmt.Println("Exibindo logs...")
	//} else if comando == 0 {
	//	fmt.Println("Tô saindo...")
	//} else {
	//	fmt.Println("Comando inválido!")
	//}

	//O operador _ (identificador em branco) significa que vou ignorar o segundo retorno dessa função
	nome, _ := devolveNomeEIdade()
	fmt.Println(nome)

	exibeNomes()

	for {
		falaOi()
		mostraMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

}

func devolveNomeEIdade() (string, int) {
	return "Gabriel", 21
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://www.google.com.br", "https://www.b3.com.br"}

	fmt.Println(sites)

	//for i := 0; i < len(sites); i++ {
	//	fmt.Println(sites[i])
	//}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("estou passando na posição", i, "e lá tem o site", site)
			testaSite(site)
		}
		time.Sleep(tempoEntreMonitoramentos * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Deu bom. O site ", site, "carregou")
	} else {
		fmt.Println("O site ", site, "está com problemas e retornou o status", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Gabriel", "Pedro", "Rocha"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Meu slice tem", len(nomes), "itens")
	fmt.Println("Esse slice tá com capacidade pra armazenar", cap(nomes), "itens")

	nomes = append(nomes, "Jurema")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Meu slice tem", len(nomes), "itens")
	fmt.Println("Esse slice tá com capacidade pra armazenar", cap(nomes), "itens")
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
	_, err := fmt.Scan(&comandoLido)
	if err != nil {
		return 0
	}
	fmt.Println("Vc escolheu o comandoLido", comandoLido)

	return comandoLido
}
