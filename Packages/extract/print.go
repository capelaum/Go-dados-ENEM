package extract

import (
	"fmt"
	"math"
)

//PrintMeanScores mostra as medias
func PrintMeanScores(medias [4]float64) {
	fmt.Println("Médias:")
	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n",
		medias[0],
		medias[1],
		medias[2],
		medias[3],
	)
}

//PrintSchoolTypeNumber mostra o numero de participantes de cada tipo de Escola
func PrintSchoolTypeNumber(schoolType [4]int) {
	fmt.Printf("\nNúmero de participantes por tipo de Escola:\n")
	fmt.Printf("\n\tNão respondeu: %d \n\tPública: %d \n\tPrivada: %d\n\tExterior: %d",
		schoolType[0],
		schoolType[1],
		schoolType[2],
		schoolType[3],
	)
}

//PrintSchoolsMeanScores mostra as medias de cada tipo de Escola
func PrintSchoolsMeanScores(schoolMeanScores [4][4]float64) {
	fmt.Printf("Médias por tipo de Escola:\n")
	fmt.Printf("Não respondeu: \n")
	PrintMeanScores(schoolMeanScores[0])
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Pública: \n")
	PrintMeanScores(schoolMeanScores[1])
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Privada: \n")
	PrintMeanScores(schoolMeanScores[2])
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Exterior: \n")
	if !math.IsNaN(schoolMeanScores[3][0]) {
		PrintMeanScores(schoolMeanScores[3])
	} else {
		fmt.Printf("Sem registros\n")
	}
}

//PrintYearData mostra os dados sobre um dado ano
func PrintYearData(year Year) {
	fmt.Printf("Dados nacionais ENEM %d", year.Year)
	fmt.Printf("\n*****************************************\n")
	PrintYearMeanScores(year)
	fmt.Printf("\n*****************************************\n")
	PrintYearRacesMeanScores(year)
	fmt.Printf("\n*****************************************\n")
	PrintYearSchoolsMeanScores(year)
	fmt.Printf("\n*****************************************\n")
	MostParticipantsUF(year.States)
	fmt.Printf("\n*****************************************\n")
	MostParticipantsRace(year.Races)
	fmt.Printf("\n*****************************************\n")
	BiggestMeanScoresState(year.States, year.Year)
	fmt.Printf("\n*****************************************\n")
	BiggestMeanScoresRace(year.Races)
	fmt.Printf("\n*****************************************\n")
	MenuStates(year.States)
}

//PrintYearMeanScores mostra as medias nacionais de um ano
func PrintYearMeanScores(year Year) {
	fmt.Printf("Médias nacionais do ano %d\n", year.Year)
	PrintMeanScores(year.Medias)
}

//PrintYearRacesMeanScores mostra as medias nacionais de raça de um ano
func PrintYearRacesMeanScores(year Year) {
	fmt.Printf("Médias nacionais por raça\n")

	for i := range year.Races {
		fmt.Printf("%s:\n", year.Races[i].Name)
		PrintMeanScores(year.Races[i].Medias)
		fmt.Printf("\n*****************************************\n")
	}
}

//PrintYearSchoolsMeanScores mostra as medias nacionais de tipo de Escola de um ano
func PrintYearSchoolsMeanScores(year Year) {
	fmt.Printf("Médias nacionais por tipo de Escola\n\n")
	PrintSchoolsMeanScores(year.SchoolMeanScores)
}

//PrintUFData printa dados de uma UF
func PrintUFData(state State) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Dados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes: %d", state.Total)
	fmt.Printf("\n*****************************************\n")
	PrintSchoolTypeNumber(state.SchoolType)
	fmt.Printf("\n*****************************************\n")
	PrintMeanScores(state.Medias)
	PrintUFRacesMeanScores(state)
	fmt.Printf("\n*****************************************\n")
	MostParticipantsRace(state.Races)
	fmt.Printf("\n*****************************************\n")
	BiggestMeanScoresRace(state.Races)
	fmt.Printf("\n*****************************************\n")
}

//PrintUFRaceMeanScores printa dados estaduais acerca de uma raça
func PrintUFRaceMeanScores(race Race) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Dados de %s", race.Name)
	fmt.Printf("\nTotal de participantes: %d\n", race.Total)
	PrintMeanScores(race.Medias)
	PrintSchoolTypeNumber(race.SchoolType)
}

//PrintUFRacesMeanScores printa dados acerca de todas raças de um estado
func PrintUFRacesMeanScores(state State) {
	for i := range state.Races {
		fmt.Printf("\n*****************************************\n")
		fmt.Printf("Dados de %s", state.Races[i].Name)
		fmt.Printf("\nTotal de participantes: %d\n", state.Races[i].Total)
		PrintMeanScores(state.Races[i].Medias)
		PrintSchoolTypeNumber(state.Races[i].SchoolType)
	}
}
