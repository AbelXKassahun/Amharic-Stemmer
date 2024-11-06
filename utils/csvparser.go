package utils

import(
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFromFile(){

}

func OpenFile(fileName string) ([][]string, error) {
	file, errReading := os.Open(fileName)
	
	if errReading != nil {
		fmt.Printf("Error opening the file <%v> \n", fileName)
		return [][]string{}, errReading
	}

	// reading the csv file
	r := csv.NewReader(file)
	record, errCSV := r.ReadAll()

	if errCSV != nil {
		fmt.Println("Error reading the CSV file ")
		return [][]string{}, errCSV
	}

	// closing the file
	errClose := file.Close()
	if errClose != nil {
		fmt.Println("Warning can't close the file: no problem tho")
	}

	return record, nil
}
