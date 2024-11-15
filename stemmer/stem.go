package stemmer

import (
	// "fmt"
	// "strings"
	"log"
	utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
)

func Stem(word string) string {
	normedInput, ok := utils.InputCheck(word, false)
	if ok != nil {
		for _, val := range ok {
			log.Println(val)
		}
		// return
		panic("invalid input") 
	}
	word = normedInput

	englishConverted := ToEng(word)


	suffixLessWord := RemoveSuffix(englishConverted)
	// fmt.Println(eng)
	amh := ToAmh(suffixLessWord)
	// fmt.Println(amh)
	return amh
}

// ʼa,bé,le/l