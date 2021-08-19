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
		exibeMenu()
		comando := leComando()
		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("Exibindo Logs")
			imprimeLogs()
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("Não conheço esse comando.")
		}
	}
}
func exibeIntroducao() {
	nome := "Saullo Almeida"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Versão: ", versao)
}
func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
	//Output:
}
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	/* sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://alura.com.br",
		"https://google.com.br",
	} */
	//for i := 0; i < len(sites); i++ {
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}
func testaSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if res.StatusCode == 200 {
		registraLog(site, true)
		fmt.Println("Site:", site, "- foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "- esta com problemas. Status Code:", res.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {
	var sites []string
	//listaSites, err := ioutil.ReadFile("lista.txt")
	arquivo, err := os.Open("lista.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			break
		}
		sites = append(sites, linha)
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
