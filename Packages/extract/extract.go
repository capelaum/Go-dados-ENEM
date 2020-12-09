package extract

import (
	"fmt"
	"io"
	"sync"
	"time"
)

/* SEM PARALELISMO */

//Data pega os dados de todos Estados do arquivo CSV
func Data(years []Year) []Year {

	now := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(now))
	}()

	fmt.Println("Extraindo dados..")

	// loop pelos anos
	for i := range years {
		reader, _ := CsvReader(years[i].CsvFilePath) // abre arquivo de cada ano
		count := 0

		// leitura de linha a linha do registro
		for /* i := 0; i < 500000; i++ */ {

			// Processa os dados linha a linha dos datasets
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} 
			if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}
			count++
			
			// Pega dados de cada estado
			for j := range years[i].States {
				if years[i].States[j].Sigla == recordLine[5] {
					years[i].States[j].Total++
					getYearSchoolScores(recordLine, &years[i])
					getUFScores(recordLine, &years[i].States[j], years[i].Year)
					getRacesData(recordLine, &years[i].States[j], years[i].Year)
				}
			}
		}

		fmt.Printf("Numero de registros analisados de %d: %d\n", years[i].Year, count)
		count = 0 // Reseta contagem a cada ano
		getStatesMeanScores(&years[i].States)
	}

	getYearsMeanScores(&years)

	return years
}

/* COM PARALELISMO */

//DataParallel pega os dados de todos Estados do arquivo CSV
func DataParallel(years *[]Year) {

	start := time.Now()
	defer func() {
		fmt.Println("\n\nTempo de execução:", time.Since(start))
	}()

	fmt.Println("Extraindo dados..")

	var wg sync.WaitGroup
	wg.Add(3)

	for i := range *years {
		reader, _ := CsvReader((*years)[i].CsvFilePath)

		totalRecords := (*years)[i].Total
		workers := 1000

		// Execução paralela dos processos de cada ano
		go func(i int) {
			defer wg.Done()

			// Numero de processos paralelos necessários para cada ano
			for j := 0; j <= workers; j++ {
				getData(
					reader,
					&(*years)[i],
					(totalRecords/workers)*j,     // inicio do pedaço
					(totalRecords/workers)*(j+1), // final do pedaço
				)
			}
		}(i)
	}
	
	wg.Wait()

	for i := range *years {
		getStatesMeanScores(&(*years)[i].States)
		fmt.Printf("Numero de registros analisados de %d: %d\n", (*years)[i].Year, (*years)[i].Total)
	}

	getYearsMeanScores(&(*years))

	return
}
