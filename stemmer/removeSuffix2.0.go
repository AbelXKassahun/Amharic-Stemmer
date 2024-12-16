package stemmer

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"strings"
	"unicode/utf8"
)

func RemoveSuffix2(word string) string {
	var suffixFound bool
	if len(word) > 3 {
		for i, val := range *utils.ReturnSuffixList() {
			if i != 0 {
				lenW := utf8.RuneCountInString(word)
				suffix := strings.ReplaceAll(strings.TrimSpace(val[0]), "|", "")
				lenS := utf8.RuneCountInString(suffix)
				if lenW <= lenS {
					continue
				}
				runedWord := []rune(word)
				lastbit := string(runedWord[lenW-lenS:])

				if strings.Contains(lastbit, suffix) {
					fmt.Printf("suffix -> %v\n", suffix)
					fmt.Printf("#lastBitOfWord -> %v\n", lastbit)
					fmt.Printf("OG word -> %v\n", word)

					// first check the rule (exception) of the suffix
					// rule is exception list and last word conversion
					// suffix removal is after rule check (exception check)
					// apply the last word conversion rule when removing suffix
					// which means after checking the exception list
					// last word conversion rule check is done outside (after) checkForExceptions function
					word = checkForExceptions(word, []rune(suffix))

					// check for last word conversion rule here
					word = word[:len(word)-len(suffix)]
					fmt.Printf("suffixless word -> %v\n", word)

					suffixFound = true
					break
				}
			}
		}
	} else { // also should be outside the loop
		return word
	}

	if !suffixFound {
		return word
	}
	return word
}
func checkForExceptions(word string, suffix []rune) string {
	// open the extension file
	// transliterate the first word we get when we split the line with -
	// check if the first word exists in word
	// if it does then transliterate the second word then replace
	// return replaced
	// no suffix removal or last word conversion rule check here
	return word
}
func convertToDiffHouse(letter string, house int) string {
	for _, val := range *utils.ReturnLetters() {
		for _, val2 := range val {
			word := strings.Split(strings.TrimSpace(val2), " ")
			if letter == word[1] {
				result := strings.Split(strings.TrimSpace(val[house-1]), " ")
				return result[1]
			}
		}
	}
	return ""
}
