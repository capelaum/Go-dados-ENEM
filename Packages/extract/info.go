package extract

import (
	"fmt"
)

/* Info Nacional */

//YearsInfo mostra info sobre os anos  
func YearsInfo(years []Year) {
	var MostParticipantsYear = years[0]
	for i := 0; i < len(years)-1; i++ {
		if years[i].Total < years[i+1].Total {
			MostParticipantsYear = years[i+1]
		}
	}
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("O Ano com maior numero de participantes: %d\n", MostParticipantsYear.Year)
	fmt.Printf("Numero de participantes: %v\n", MostParticipantsYear.Total)

	total := 0.0
	medias := [4]float64{}

	for i := range years {
		for j := range years[i].Medias {
			total += years[i].Medias[j]
		}
		medias[i] = total/4
		total = 0.0
	}

	biggestMeanScoresYear := Year{}

	for i := 0; i < len(medias)-1; i++ {
		if medias[i] < medias[i+1] {
			biggestMeanScoresYear = years[i+1]
		}
	}
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("O Ano com as maiores medias: %d\n", biggestMeanScoresYear.Year)
	PrintMeanScores(biggestMeanScoresYear.Medias)
}

//MostParticipantsUF mostra o estado com maior numero de participantes de um ano
func MostParticipantsUF(states []State) {
	var maior = states[0]
	for i := 0; i < len(states)-1; i++ {
		if states[i].Total < states[i+1].Total {
			maior = states[i+1]
		}
	}
	fmt.Printf("O Estado com maior numero de participantes: %s\n", maior.Sigla)
	fmt.Printf("Numero de participantes: %v", maior.Total)
}

//MostParticipantsRace mostra a raça com maior numero de participantes de um ano
func MostParticipantsRace(races [6]Race) {
	var MostParticipantsRace = races[0]

	for i := 0; i < len(races)-1; i++ {
		if races[i].Total < races[i+1].Total {
			MostParticipantsRace = races[i+1]
		}
	}
	fmt.Printf("A Raça com maior numero de participantes: %s\n", MostParticipantsRace.Name)
	fmt.Printf("Numero de participantes: %v\n", MostParticipantsRace.Total)
}

//BiggestMeanScoresState mostra o estado com as maiores media de um dado ano
func BiggestMeanScoresState(states []State, year int) {
	total := 0.0
	medias := [27]float64{}

	for i := range states {
		for j := range states[i].Medias {
			total += states[i].Medias[j]
		}
		medias[i] = total/4
		total = 0.0
	}

	biggestMeanScoresState := State{}

	for i := 0; i < len(medias)-1; i++ {
		if medias[i] < medias[i+1] {
			biggestMeanScoresState = states[i+1]
		}
	}

	fmt.Printf("O Estado com as maiores médias do ano %d: %s\n", year, biggestMeanScoresState.Sigla)
	PrintMeanScores(biggestMeanScoresState.Medias)
}

//BiggestMeanScoresRace mostra a raça com as maiores médias 
func BiggestMeanScoresRace(races [6]Race) {
	total := 0.0
	medias := [6]float64{}

	for i := range races {
		for j := range races[i].Medias {
			total += races[i].Medias[j]
		}
		medias[i] = total/4
		total = 0.0
	}

	biggestMeanScoresRace := Race{}

	for i := 0; i < len(medias)-1; i++ {
		if medias[i] < medias[i+1] {
			biggestMeanScoresRace = races[i+1]
		}
	}

	fmt.Printf("A Raça com as maiores médias: %s\n", biggestMeanScoresRace.Name)
	PrintMeanScores(biggestMeanScoresRace.Medias)
}
