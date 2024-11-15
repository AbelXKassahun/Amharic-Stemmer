package utils

import(
	"encoding/csv"
	"path/filepath"
	"fmt"
	"os"
)

func ReadFromFile(){

}

func OpenFile(fileName string) ([][]string, error) {
	var errorMsg error
	file, errReading := os.Open(fileName)
	
	if errReading != nil {
		errorMsg = fmt.Errorf("error opening the file <%v> \n", fileName)
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
	absPath, _ := filepath.Abs("../data/letters.csv")
	// record, err := OpenFile("../data/letters.csv")
	record, err := OpenFile(absPath)
	if err != nil {
		panic(err)
	}
	return &record
}

func ReturnSuffixList() *[][]string {
	absPath, _ := filepath.Abs("../data/suffixList.csv")
	// suffixList, err := OpenFile("../data/suffixList.csv")
	suffixList, err := OpenFile(absPath)
	if err != nil {
		panic(err)
	}
	return &suffixList
}