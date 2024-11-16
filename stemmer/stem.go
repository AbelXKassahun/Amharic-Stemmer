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
		var errorMsgs string
		for _, val := range ok {
			errorMsgs += val.Error() + "\n"
		}
		log.Println("from initial input check")
		log.Println(ok)
		return errorMsgs
	}
	word = normedInput

	englishConverted, err := ToEng(word)
	if err != nil {
		log.Println("From ToEng ----")
		log.Println(err.Error())
		return err.Error()
	}


	suffixLessWord := RemoveSuffix(englishConverted)
	// fmt.Println(eng)
	amh, err := ToAmh(suffixLessWord)
	if err != nil {
		log.Println("From ToAmh ----")
		log.Println(err.Error())
		return err.Error()
	}
	// fmt.Println(amh)
	return amh
}

// ʼa,bé,le/l