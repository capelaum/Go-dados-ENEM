package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/montanaflynn/stats"
)

// getData pega os dados de forma paralela, lendo uma quantidade de linhas do arquivo
func getData(
	reader *csv.Reader,
	year *Year,
	start int,
	end int,
) {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// leitura de pedaços do csv paralelamente
		for i := start; i < end; i++ {
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} 
			if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}
			// Processa os dados do arquivo linha alinha de forma paralela
			for i := range (*year).States {
				if (*year).States[i].Sigla == recordLine[5] {
					(*year).States[i].Total++
					getYearSchoolScores(recordLine, &(*year))
					getUFScores(recordLine, &(*year).States[i], year.Year)
					getRacesData(recordLine, &(*year).States[i], year.Year)
				}
			}
		}
	}()
	wg.Wait()
}

// Pega as notas de areas de conhecimento de cada UF
func getUFScores(recordLine []string, state *State, year int) {
	getScores(recordLine, &state.Scores, year)
}

func getScores(recordLine []string, scores *[4][]float64, year int) {
	// Ano 2019
	if year == 2019 {
		scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
		scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
		scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
		scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94
	}

	// Anos 2017 e 2018
	scores[0] = append(scores[0], getScore(recordLine, 90)) // Ciencias da Natureza - campo 90
	scores[1] = append(scores[1], getScore(recordLine, 91)) // Ciencias Humanas - campo 91
	scores[2] = append(scores[2], getScore(recordLine, 92)) // Linguagens e Código - campo 92
	scores[3] = append(scores[3], getScore(recordLine, 93)) // Matemática - campo 93
}

func getScore(recordLine []string, campo int) float64 {
	i, _ := strconv.ParseFloat(recordLine[campo], 64)
	return i
}

func getIntValue(recordLine []string, campo int) int {
	i, _ := strconv.Atoi(recordLine[campo])
	return i
}

// Pega os dados de cada tipo de Raça
func getRacesData(recordLine []string, s *State, year int) {

	// Tipo de cor/raça - campo 9
	raceType := getIntValue(recordLine, 9)

	// Tipo de escola - campo 17
	schoolType := getIntValue(recordLine, 17)

	switch raceType {
	case 0: // Não informado
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 1: // Branca
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 2: // Preta
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 3: // Parda
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 4: // Amarela
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	case 5: // Indigena
		getRaceTypeData(recordLine, s, raceType, schoolType, year)
	default:
		fmt.Println("Raça não reconhecida, possível E.T!")
	}
}

// Pega os dados de uma Raça
func getRaceTypeData(recordLine []string, s *State, raceType int, schoolType int, year int) {
	s.Races[raceType].Total++
	getRaceScores(recordLine, s, raceType, year)
	getSchoolType(s, raceType, schoolType)
}

// Pega as notas de areas de conhecimento de Cada Raça
func getRaceScores(recordLine []string, state *State, raceType int, year int) {
	getScores(recordLine, &state.Races[raceType].Scores, year)
}

// pega quantidade de participantes de cada tipo de escola por estado e raça
func getSchoolType(s *State, raceType int, schoolType int) {
	switch schoolType {
	case 1: // Nao respondeu
		s.SchoolType[0]++
		s.Races[raceType].SchoolType[0]++
	case 2: // Publica
		s.SchoolType[1]++
		s.Races[raceType].SchoolType[1]++
	case 3: // Privada
		s.SchoolType[2]++
		s.Races[raceType].SchoolType[2]++
	case 4: // Exterior
		s.SchoolType[3]++
		s.Races[raceType].SchoolType[3]++
	default:
		fmt.Println("Algo errado, tipo de escola invalido!")
	}
}

// Calcula as medias das notas de cada area de conhecimento
func getMeanScores(scores [4][]float64) [4]float64 {
	meanScores := [4]float64{}
	
	for i := range scores {
		meanScores[i], _ = stats.Mean(scores[i])
	}
	return meanScores
}

// Pega as medias das notas de cada area de conhecimento e por raça de todos Estados
func getStatesMeanScores(states *[]State) {
	for i := range *states {
		(*states)[i].Medias = getMeanScores((*states)[i].Scores)

		for j := range (*states)[i].Races {
			(*states)[i].Races[j].Medias = getMeanScores((*states)[i].Races[j].Scores)
		}
	}
}

// Pega as medias nacionais de cada ano
func getYearsMeanScores(years *[]Year) {
	statesMeanScores := [4][]float64{}

	for i := range *years {
		for j := range (*years)[i].States {

			for k := range (*years)[i].States[j].Medias {
				// pega todas 4 medias dos 26 estados de cada ano
				statesMeanScores[k] = append(statesMeanScores[k], (*years)[i].States[j].Medias[k])
			}

			for k := range (*years)[i].States[j].Races {
				(*years)[i].Races[k].Total += (*years)[i].States[j].Races[k].Total

				for l := range (*years)[i].States[j].Races[k].Medias {
					// pega todas 4 medias de cada uma das 6 raças de cada um dos 26 estados de cada ano
					(*years)[i].Races[k].Scores[l] =
						append((*years)[i].Races[k].Scores[l], (*years)[i].States[j].Races[k].Medias[l])
				}
			}
		}
		(*years)[i].Medias = getMeanScores(statesMeanScores)

		// tira as medias das 4 medias de cada raça dos 26 estados de cada ano
		for k := range (*years)[i].Races { 
			(*years)[i].Races[k].Medias = getMeanScores((*years)[i].Races[k].Scores)
		}

		for j := range (*years)[i].SchoolMeanScores {
			(*years)[i].SchoolMeanScores[j] = getMeanScores((*years)[i].SchoolScores[j])
		}
	}
}

// Pega os Scores nacionais de cada ano por tipo de escola do ensino medio
func getYearSchoolScores(recordLine []string, year *Year) {
	schoolType := getIntValue(recordLine, 17) // Tipo de Escola - Campo 17

	switch schoolType {
	case 1: // Nao respondeu
		getScores(recordLine, &year.SchoolScores[0], year.Year)
	case 2: // Publica
		getScores(recordLine, &year.SchoolScores[1], year.Year)
	case 3: // Privada
		getScores(recordLine, &year.SchoolScores[2], year.Year)
	case 4: // Exterior
		getScores(recordLine, &year.SchoolScores[3], year.Year)
	default:
		fmt.Println("Algo errado, tipo de escola invalido!")
	}
}
