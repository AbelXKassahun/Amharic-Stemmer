package utils

import(
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFromFile(){

}

func OpenFile(fileName string) ([][]string, error) {
	var errorMsg error
	file, errReading := os.Open(fileName)
	
	if errReading != nil {
		errorMsg = fmt.Errorf("Error opening the file <%v> \n", fileName)
		return [][]string{}, errorMsg
	}

	// reading the csv file
	r := csv.NewReader(file)
	record, errCSV := r.ReadAll()

	if errCSV != nil {
		errorMsg = fmt.Errorf("Error reading the CSV file: %v", fileName)
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
	record, err := OpenFile("../data/letters.csv")
	if err != nil {
		panic(err)
	}
	return &record
}

func ReturnSuffixList() *[][]string {
	suffixList, err := OpenFile("../data/suffixList.csv")
	if err != nil {
		panic(err)
	}
	return &suffixList
}