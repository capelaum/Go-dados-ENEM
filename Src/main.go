package main

import (
	"fmt"
	"os"
	"runtime"

	"../Packages/extract"
)

// Definido para utiliza o maximo de processadores  
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	years := extract.NewYears()

	var opcao int
	fmt.Println("Escolha uma opção para extrair dados:")
	fmt.Println("Digite -1 para sair")
	fmt.Println("\n1) Sem Paralelismo")
	fmt.Println("2) Com Paralelismo")
	fmt.Print("\n-> ")
	fmt.Scan(&opcao)

	switch opcao {
	case -1:
		os.Exit(3)
	case 1:
		years = extract.Data(years)
	case 2:
		extract.DataParallel(&years)
	default:
		fmt.Println("Opção Inválida!")
		os.Exit(3)
	}
	// Apresentação dos dados anuais e Menus
	extract.YearsInfo(years)
	extract.YearsMenu(years)

	return
}
