package extract

import (
	"fmt"
	"os"
)

//YearsMenu - para escolher quais dados mostrar sobre anos
func YearsMenu(years []Year) {
	for {
		fmt.Printf("\n*****************************************\n")
		fmt.Println("Escolha o ano.")
		fmt.Printf("Digite -1 para sair\n\n")
		fmt.Printf("0-Enem 2017\n1-Enem 2018\n2-Enem 2019\n")
		fmt.Print("\n-> ")

		var year int
		fmt.Scan(&year)

		fmt.Printf("\n*****************************************\n")
		// verifica se existe UF no arrays de structs states
		switch year {
		case -1:
			os.Exit(3)
		case 0:
			PrintYearData(years[year])
		case 1:
			PrintYearData(years[year])
		case 2:
			PrintYearData(years[year])
		default:
			fmt.Println("Opção inválida!")
			os.Exit(3)
		}
		fmt.Printf("\n*****************************************\n")
	}
}

//MenuStates para escolher quais dados mostrar sobre estados
func MenuStates(states []State) {
	for {
		fmt.Println("Escolha de qual UF deseja visualizar dados: ")
		fmt.Printf("Digite -1 para sair\n\n")
		for i := range states {
			fmt.Printf("%s ", states[i].Sigla)
		}
		fmt.Print("\n-> ")

		var UF string
		fmt.Scan(&UF)

		// verifica se existe UF no arrays de structs states
		if Contains(states, UF) == true {
			for i := range states {
				if UF == states[i].Sigla {
					PrintUFData(states[i])
					MenuRaces(states[i])
				}
			}
		} 

		if UF == "-1" {
			break
		}

		fmt.Println("Opção inválida.")
		UF = "" // reseta UF digitada pelo usuario
	}
}

//MenuRaces para escolher quais dados mostrar sobre raças
func MenuRaces(state State) {
	fmt.Printf("\n\nDeseja visualizar dados por Raça em %s? (s/n)", state.Sigla)
	fmt.Print("\n-> ")
	var opcao string
	fmt.Scan(&opcao)

	if opcao == "n" {
		return
	}
	for {
		fmt.Printf("\n*****************************************\n")
		fmt.Println("Escolha a raça.")
		fmt.Println("Digite -1 para sair.")
		fmt.Printf("\n0-Todas Raças\n1-Não informada\n2-Branca\n3-Preta\n4-Parda\n5-Amarela\n6-Indígena\n")
		fmt.Print("\n-> ")

		var raca int
		fmt.Scan(&raca)

		switch raca {
		case 0:
			PrintUFRacesMeanScores(state)
		case 1:
			PrintUFRaceMeanScores(state.Races[raca-1])
		case 2:
			PrintUFRaceMeanScores(state.Races[raca-1])
		case 3:
			PrintUFRaceMeanScores(state.Races[raca-1])
		case 4:
			PrintUFRaceMeanScores(state.Races[raca-1])
		case 5:
			PrintUFRaceMeanScores(state.Races[raca-1])
		case 6:
			PrintUFRaceMeanScores(state.Races[raca-1])
		}

		if raca == -1 {
			break
		}
		if raca > 6 {
			fmt.Println("Opção inválida.")
			break
		}
	}
}
