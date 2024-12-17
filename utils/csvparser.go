package utils

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "path/filepath"
)

//go:embed data/letters.csv
var lettersCSV []byte

//go:embed data/suffixList.csv
var suffixesCSV []byte

//go:embed data/prefixList.csv
var prefixesCSV []byte

//go:embed data/exceptions.csv
var exceptionsCSV []byte

func OpenFile(fileName string) ([][]string, error) {
	var errorMsg error
	file, errReading := os.Open(fileName)

	if errReading != nil {
		errorMsg = fmt.Errorf("error opening the file <%v>", fileName)
		return [][]string{}, errorMsg
	}

	// reading the csv file
	r := csv.NewReader(file)
	record, errCSV := r.ReadAll()

	if errCSV != nil {
		errorMsg = fmt.Errorf("error reading the CSV file: %v", fileName)
		return [][]string{}, errorMsg
	}

	// closing the file
	errClose := file.Close()
	if errClose != nil {
		fmt.Println("Warning can't close the file: no problem tho")
	}

	return record, nil
}

func ReturnLetters() *[][]string {
	// absPath, _ := filepath.Abs("../data/letters.csv")
	// record, err := OpenFile(absPath)
	// log.Println(absPath)
	// if err != nil {
	// 	panic(err)
	// }
	reader := csv.NewReader(bytes.NewReader(lettersCSV))
	record, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return &record
}

func ReturnSuffixList() *[][]string {
	// absPath, _ := filepath.Abs("../data/suffixList.csv")
	// suffixList, err := OpenFile(absPath)
	// if err != nil {
	// 	panic(err)
	// }
	reader := csv.NewReader(bytes.NewReader(suffixesCSV))
	suffixList, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return &suffixList
}

func ReturnPrefixList() *[][]string {
	reader := csv.NewReader(bytes.NewReader(prefixesCSV))
	prefixList, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return &prefixList
}

func ReturnExceptionsList() *[][]string {
	reader := csv.NewReader(bytes.NewReader(exceptionsCSV))
	exceptionsList, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return &exceptionsList
}
