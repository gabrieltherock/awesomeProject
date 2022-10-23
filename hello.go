package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
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

	leSitesDoArquivo()

	for {
		falaOi()
		mostraMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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

	//sites := []string{"https://www.alura.com.br", "https://www.google.com.br", "https://www.b3.com.br"}
	sites := leSitesDoArquivo()

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

func leSitesDoArquivo() []string {
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(file)
	var sites []string

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}
	}

	_ = file.Close()
	fmt.Println(sites)
	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Deu bom. O site ", site, "carregou")
		registraLog(site, true)
	} else {
		fmt.Println("O site ", site, "está com problemas e retornou o status", resp.StatusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, b bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	_, _ = arquivo.WriteString(time.Now().Format("2006/01/02 15:04:05 ") + site + " - online: " + strconv.FormatBool(b) + "\n")

	_ = arquivo.Close()
	fmt.Println(arquivo)
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
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
