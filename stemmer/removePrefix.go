package stemmer

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"regexp"
	"strings"
	"unicode/utf8"
)

func RemovePrefix(word string) string {
	lenW := utf8.RuneCountInString(word)
	if lenW <= 3 {
		return word
	}

	for i, val := range *utils.ReturnPrefixList() {
		if i != 0 {
			prefix := strings.ReplaceAll(strings.TrimSpace(val[0]), "|", "")
			lenP := utf8.RuneCountInString(prefix)
			if lenP >= lenW {
				continue
			}

			runedWord := []rune(word)
			firstBitOfWord := string(runedWord[:lenP])

			if strings.Contains(firstBitOfWord, prefix) {
				fmt.Printf("first bit of word -> %v\n", firstBitOfWord)
				fmt.Printf("prefix -> %v\n", prefix)
				// all if's will be refactored to a function (checkException function)
				if prefix == "la" {
					return laPrefixHandler(word)
				} else {
					// blatant removal
					word = string(runedWord[lenP:])
					fmt.Printf("prefixless word -> %v\n", word)

					return word
				}
			}
		}
	}
	return word
}

// dealing with the prefix ለ
// this would only work if
func laPrefixHandler(word string) string {
	var exceptions = []string{"lanaga", "laqay"}

	// if the suffix remover works perfectly replace ([aā]?) with a better alternative like ([a])
	// you need to reconsider the last capture group based on your progress in vowel starting suffix removal
	pattern := `^(l)(a)([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpv]|ǧ|p̣)([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpv]|ǧ|p̣)([aā]?)$`
	re := regexp.MustCompile(pattern)

	for _, val := range exceptions {
		if strings.Contains(word, val) {
			return strings.Replace(word, "la", "", 1)
		}
	}

	if !re.MatchString(word) {
		return strings.Replace(word, "la", "", 1)
	} else {
		fmt.Println("maatches ahhhh")
	}
	//if !re.MatchString(word) {
	//	match := re.FindString(word)
	//	fmt.Println("matches ", match)
	//	return strings.Replace(word, "la", "", 1)
	//}

	return word
}
