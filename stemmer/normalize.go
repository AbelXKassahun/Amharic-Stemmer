package stemmer

import (
	"regexp"
	"strings"
)

func inputCheck(word string) string {
	word = strings.ReplaceAll(word, " ", "")

	// pattern := `^(?=.*[A-Za-z])(?=.*\d)(?=.*[^\w\s]).+$`
	alphabets := `([A-Za-z]+)`
	// ፩፪፫፬፭፮፮፰፱፲፳፴፵፵፷፸፹፺፻
	digits := `([፩-፻0-9]+)`
	// :፤። [^\p{L}\s]
	// symbols := `([.\?,/#!$@%^&*;:\+፤።{}\.\[\]=\-_~()]+)`
	symbols := `[^\p{L}\s]`

	isLetter := regexp.MustCompile(alphabets)
	isDigit := regexp.MustCompile(digits)
	isSymbol := regexp.MustCompile(symbols)

	if isLetter.MatchString(word) {
		word = isLetter.ReplaceAllString(word, "")
	}
	if isDigit.MatchString(word) {
		word = isDigit.ReplaceAllString(word, "")
	}
	if isSymbol.MatchString(word) {
		word = isSymbol.ReplaceAllString(word, "")
	}

	return word
}

func HandleRedundantLetters(word string) string {
	redundant := []string{
		"ዐ አ", "ዑ ኡ", "ዒ ኢ", "ዓ ኣ", "ዔ ኤ", "ዕ እ", "ዖ ኦ",
		"ሐ ሀ", "ሑ ሁ", "ሒ ሂ", "ሓ ሃ", "ሔ ሄ", "ሕ ህ", "ሖ ሆ",
		"ኀ ሀ", "ኁ ሁ", "ኂ ሂ", "ኃ ሃ", "ኄ ሄ", "ኅ ህ", "ኆ ሆ",
		"ሠ ሰ", "ሡ ሱ", "ሢ ሲ", "ሣ ሳ", "ሤ ሴ", "ሥ ስ", "ሦ ሶ", // x
		"ጸ ፀ", "ጹ ፁ", "ጺ ፂ", "ጻ ፃ", "ጼ ፄ", "ጽ ፅ", "ጾ ፆ", // x
	}

	for _, val := range redundant {
		words := strings.Split(val, " ")
		if strings.Contains(word, words[0]) {
			word = strings.ReplaceAll(word, words[0], words[1])
		}
	}

	return word
}

//func HandleAbreviations(word string) string {
//
//	return word
//}
