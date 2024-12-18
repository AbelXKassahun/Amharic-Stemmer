package stemmer

import (
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
				aFamily := []string{"at", "an", "ay", "al"}
				//fmt.Printf("first bit of word -> %v\n", firstBitOfWord)
				//fmt.Printf("prefix -> %v\n", prefix)
				// all if's will be refactored to a function (checkException function)
				for _, val := range aFamily {
					if prefix == val {
						return aPrefixHandler(word)
					}
				}
				if prefix == "la" {
					return laPrefixHandler(word)
				} else if prefix == "t" {
					return tPrefixHandler(word)
				} else if prefix == "ka" {
					return KaHandler(word)
				} else {
					// blatant removal
					word = string(runedWord[lenP:])
					//fmt.Printf("prefixless word -> %v\n", word)

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
	pattern := `^(l)(a)([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([aā]?)$`
	re := regexp.MustCompile(pattern)

	for _, val := range exceptions {
		if strings.Contains(word, val) {
			return strings.Replace(val, "la", "", 1)
		}
	}

	if !re.MatchString(word) {
		return strings.Replace(word, "la", "", 1)
	}
	//if !re.MatchString(word) {
	//	match := re.FindString(word)
	//	fmt.Println("matches ", match)
	//	return strings.Replace(word, "la", "", 1)
	//}

	return word
}

func aPrefixHandler(word string) string {
	// a, n, C, a|o, C, ā?, C?, C
	pattern := `^([a])([tnyls])(.{2,})([mwn])$`
	re := regexp.MustCompile(pattern)

	if re.MatchString(word) {
		//return strings.Replace(word, "an", "", 1)
		return word[2:]
	} else {
		pattern := `^([a])([tnyls])(.{2,})([ā])$`
		exclude := "bas"
		re := regexp.MustCompile(pattern)
		if re.MatchString(word) && !containsExcludedSequence(word, exclude) {
			return word[2:]
		}
	}

	return word
}

func containsExcludedSequence(input, exclude string) bool {
	return regexp.MustCompile(exclude).MatchString(input)
}

func tPrefixHandler(word string) string {
	firstPattern := `^([t])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])`
	secondPattern := `^([t])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ā])`
	pattern1 := regexp.MustCompile(firstPattern)
	pattern2 := regexp.MustCompile(secondPattern)

	if pattern1.MatchString(word) {
		word = pattern1.ReplaceAllStringFunc(word, func(match string) string {
			subMatches := pattern1.FindStringSubmatch(match)
			return subMatches[2] + subMatches[3]
		})
	} else if pattern2.MatchString(word) {
		word = pattern2.ReplaceAllStringFunc(word, func(match string) string {
			subMatches := pattern2.FindStringSubmatch(match)
			return subMatches[2] + subMatches[3] + subMatches[4] + subMatches[5] + subMatches[6]
		})
	}
	return word
}

func KaHandler(word string) string {
	exceptions := []string{"kadamoz-damoz"}
	pattern := `^([k])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([no])`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(word) {
		word = word[2:]
	} else {
		for _, val := range exceptions {
			exp := strings.Split(val, "-")
			if strings.Contains(word, exp[0]) {
				word = strings.Replace(word, exp[0], exp[1], 1)
			}
		}
	}

	return word
}
