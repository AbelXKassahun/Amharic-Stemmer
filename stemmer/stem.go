package stemmer

import (
	"fmt"
	"strings"
)

func Stem(record [][]string, word string) []string {
	var engWord []string
	letters := strings.Split(word, "")
	
	for _, val := range letters {
		if result, ok := LookForLetter(record, val); ok != nil{
			fmt.Printf("\n")
			panic(ok)
		} else {
			engWord = append(engWord, result)
			fmt.Println(result)
		}
	}

	return engWord
}