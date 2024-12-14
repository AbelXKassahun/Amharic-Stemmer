package stemmer

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"strings"
	"unicode/utf8"
)

// RemoveSuffix - removes suffix with the longest match
func RemoveSuffix(word string) []string {
	var suffixless []string
	for i, val := range *utils.ReturnSuffixList() {
		if i != 0 {
			lenW := utf8.RuneCountInString(word)
			if len(word) > 3 {
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
					// lastbit = strings.ReplaceAll(lastbit, suffix, "")
					fmt.Printf("OG word -> %v\n", word)
					word = word[:len(word)-len(suffix)]
					fmt.Printf("suffixless word -> %v\n", word)
					// word += lastbit
					suffixless = append(suffixless, word)
					// check if the first token of the suffix is a vowel
					// if it is a vowel change the house (to 4th house) of the consonant that lost its vowel pair
					handled := CheckForExceptions(word, []rune(suffix))
					if handled != "" {
						fmt.Printf("handled -> %v\n", handled)
						suffixless = append(suffixless, handled)
					} else {
						fmt.Printf("no-handled -> %v\n", handled)
					}
					break
				}
			} else {
				suffixless = append(suffixless, word)
			}
		}
	}
	return suffixless
}

func CheckForExceptions(word string, suffix []rune) string {
	vowels := []string{"a", "u", "i", "ā", "é", "e", "o"}
	var handled string
	firstLetterOfSuffix := string(suffix[0])

	if firstLetterOfSuffix == vowels[3] || firstLetterOfSuffix == vowels[6] {
		fmt.Printf("suffix[0] -> %v\n", firstLetterOfSuffix)
		handled = exception1(word)
	}
	// multiple exception could exist on the same string (so no else if )
	// so chain them (pass the handled to the next exception check)

	return handled
}

func exception1(word string) string {
	var handled string
	lenW := utf8.RuneCountInString(word)
	runedWord := []rune(word)
	lastLetter := string(runedWord[lenW-1])
	fmt.Printf("lastletter -> %v\n", lastLetter)
	fourthLetter := convertToDiffHouse(lastLetter, 4)
	fmt.Printf("fourthletter -> %v\n", fourthLetter)
	//handled = word[:len(word)-1] + fourthLetter
	handled = string(runedWord[:lenW-1]) + fourthLetter
	return handled
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
