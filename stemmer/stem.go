package stemmer

import (
	// "fmt"
	// "strings"
	// utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
)

func Stem(word string) string {
	englishConverted := ToEng(word)


	suffixLessWord := RemoveSuffix(englishConverted)
	// fmt.Println(eng)
	amh := ToAmh(suffixLessWord)
	// fmt.Println(amh)
	return amh
}

// ʼa,bé,le/l