package stemmer

import (
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
	// "fmt"
	// "strings"
	"log"
)

func Stem(word string) ([]string, string) {
	normedInput, ok := utils.InputCheck(word, false)
	if ok != nil {
		var errorMsgs string
		for _, val := range ok {
			errorMsgs += val.Error() + "\n"
		}
		log.Println("from initial input check")
		log.Println(ok)
		return []string{}, errorMsgs
	}
	word = normedInput

	englishConverted, err := ToEng(word)
	if err != nil {
		log.Println("From ToEng ----")
		log.Println(err.Error())
		return []string{}, err.Error()
	}

	prefixLessWord := RemovePrefix(englishConverted)
	infixLessWord := RemoveInfix(prefixLessWord)
	suffixLessWord := RemoveSuffix(infixLessWord)
	if len(suffixLessWord) != 0 { // only remove vowels when the word has affixes
		suffixLessWord = append(suffixLessWord, removeVowels(suffixLessWord[0]))
	}

	var affixless []string
	for _, val := range suffixLessWord {
		amh, err := ToAmh(val)
		if err != nil {
			log.Println("From ToAmh ----")
			log.Println(err.Error())
			return []string{}, err.Error()
		}
		affixless = append(affixless, amh)
	}
	// fmt.Println(amh)
	return affixless, ""
}

// ʼa,bé,le/l
