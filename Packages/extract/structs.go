package extract

//Year - Estrutura de um ano do Enem
type Year struct {
	Year             int
	States           []State
	Total            int
	Medias           [4]float64
	Races            [6]Race
	SchoolScores     [4][4][]float64 // 4 tipos de Escola, 4 areas, varias notas
	SchoolMeanScores [4][4]float64   // 4 tipos de Escola, 4 medias

	CsvFilePath string
}

//State - Estrutura de Estado (UF)
type State struct {
	Sigla      string
	Total      int
	Medias     [4]float64
	Scores     [4][]float64 // 4 areas de conhecimento
	Races      [6]Race      // 6 raças ao total
	SchoolType [4]int       // quantidade de registros dos 4 tipos de Escola
}

//Race - Estrutura de Raças
type Race struct {
	Name       string
	RaceType   int
	Total      int
	Scores     [4][]float64
	Medias     [4]float64
	SchoolType [4]int
}

//NewYears construtor de array de Anos do Enem
func NewYears() []Year {

	years := []Year{}
	states17 := NewStates()
	states18 := NewStates()
	states19 := NewStates()
	races17 := NewRaces()
	races18 := NewRaces()
	races19 := NewRaces()
	csvFilePath17 := "../microdados_enem_2017/DADOS/MICRODADOS_ENEM_2017.csv"
	csvFilePath18 := "../microdados_enem_2018/DADOS/MICRODADOS_ENEM_2018.csv"
	csvFilePath19 := "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"

	year17 := Year{
		Year:        2017,
		States:      states17,
		CsvFilePath: csvFilePath17,
		Total:       6731342, // total de registros
		Races:       races17,
	}

	year18 := Year{
		Year:        2018,
		States:      states18,
		CsvFilePath: csvFilePath18,
		Total:       5513748,
		Races:       races18,
	}

	year19 := Year{
		Year:        2019,
		States:      states19,
		CsvFilePath: csvFilePath19,
		Total:       5095271,
		Races:       races19,
	}

	years = append(years, year17)
	years = append(years, year18)
	years = append(years, year19)

	return years
}

//NewStates construtor de array de Estruturas de Estado (UF)
func NewStates() []State {
	siglas := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	states := []State{}
	races := NewRaces()

	for i := range siglas {
		state := State{
			Races: races,
		}
		states = append(states, state)
		states[i].Sigla = siglas[i]
		states[i].Total = 0
	}
	return states
}

//NewRaces construtor de array de Raças
func NewRaces() [6]Race {
	races := [6]Race{
		Race{Name: "Raça Não Informada", RaceType: 0},
		Race{Name: "Raça Branca", RaceType: 1},
		Race{Name: "Raça Preta", RaceType: 2},
		Race{Name: "Raça Parda", RaceType: 3},
		Race{Name: "Raça Amarela", RaceType: 4},
		Race{Name: "Raça Indigena", RaceType: 5},
	}
	return races
}
