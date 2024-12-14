package stemmer

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
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

				word = string(runedWord[lenP:])
				fmt.Printf("prefixless word -> %v\n", word)

				return word
			}
		}
	}
	return word
}
