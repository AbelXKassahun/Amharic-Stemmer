package stemmer

import (
	// "fmt"
	// "strings"
	"log"
)

func Stem(word string) ([]string, string) {
	word = inputCheck(word)
	// check if the word is a stop word here.
	word = HandleRedundantLetters(word)

	englishConverted, err := ToEng(word)
	if err != nil {
		log.Println("From ToEng ----")
		log.Println(err.Error())
		return []string{}, err.Error()
	}

	prefixLessWord := RemovePrefix(englishConverted)
	suffixLessWord := RemoveSuffix2(prefixLessWord)
	infixLessWord := RemoveInfix(suffixLessWord)
	result := []string{infixLessWord, removeVowels(infixLessWord)}
	//suffixLessWord := RemoveSuffix(infixLessWord)
	//if len(suffixLessWord) != 0 { // only remove vowels when the word has affixes
	//	suffixLessWord = append(suffixLessWord, removeVowels(suffixLessWord[0]))
	//}
	var affixless []string
	for _, val := range result {
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
