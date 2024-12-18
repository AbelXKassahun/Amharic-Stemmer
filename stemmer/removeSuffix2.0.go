package stemmer

import (
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func RemoveSuffix2(word string) string {
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

				if strings.Contains(lastbit, suffix) { // if lastbit == suffix {
					//fmt.Printf("suffix -> %v\n", suffix)
					//fmt.Printf("#lastBitOfWord -> %v\n", lastbit)
					//fmt.Printf("OG word -> %v\n", word)
					var rule []string
					if strings.ReplaceAll(val[1], " ", "") != "" {
						rule = strings.Split(strings.TrimSpace(val[1]), "|")
					} else {
						rule = []string{"#", "0"}
					}

					if suffix == "n" {
						result, found := nHandler(word)
						if found {
							word = result
							break
						}
					}

					//fmt.Printf("val:%v, rule:%v, len:%v, suffix:%v \n", val[1], rule, len(rule), suffix)
					word = word[:len(word)-len(suffix)]
					//fmt.Printf("suffixless word -> %v\n", word)
					if len(rule) != 0 {
						if rule[0] != "#" { // the suffix has an exception list
							var found bool
							word, found = checkForExceptions(word, rule)
							if !found {
								lwcRule, _ := strconv.Atoi(rule[1])
								if lwcRule != 0 && lwcRule != -1 { // convert the last letter after the suffix removal
									word = lastWordConversion(word, lwcRule)
								} else if lwcRule == -1 {
									word += suffix // add the suffix back lmfao
								}
							}
						} else { // suffix has no exception list
							lwcRule, _ := strconv.Atoi(rule[1])
							if lwcRule != 0 { // convert the last letter after the suffix removal
								word = lastWordConversion(word, lwcRule)
							}
						}
					}
					break
				}
			}
		}
	}
	return word
}

func nHandler(word string) (string, bool) {
	var found bool
	firstPattern := `([ā])([n])$`
	secondPattern := `([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([n])$`
	pattern1 := regexp.MustCompile(firstPattern)
	pattern2 := regexp.MustCompile(secondPattern)

	if pattern1.MatchString(word) {
		word = pattern1.ReplaceAllStringFunc(word, func(match string) string {
			subMatches := pattern1.FindStringSubmatch(match)
			return subMatches[1]
		})
		found = true
	} else if pattern2.MatchString(word) {
		word = pattern2.ReplaceAllStringFunc(word, func(match string) string {
			subMatches := pattern2.FindStringSubmatch(match)
			return subMatches[1] + "a"
		})
		found = true
	}
	return word, found
}

func lastWordConversion(word string, lwcRule int) string {
	lastLetter := word[len(word)-1:]
	newLastLetter := convertToDiffHouse(lastLetter, lwcRule)
	word = word[:len(word)-1]
	word += newLastLetter
	return word
}

func checkForExceptions(word string, rule []string) (string, bool) {
	// open the extension file
	lineStart, _ := strconv.Atoi(rule[0])
	lineEnd := "####"

	for i, val := range *utils.ReturnExceptionsList() {
		if i == 0 {
			continue
		}
		if i >= lineStart-1 && val[0] != lineEnd {
			// transliterate the first word we get when we split the line with -
			exceptionRow := strings.Split(val[0], "-")
			exception, err := ToEng(exceptionRow[0])
			if err != nil {
				log.Println("From checkForExceptions (exceptions) ----")
				log.Println(err.Error())
				panic(err)
			}
			if strings.Contains(word, exception) {
				newWord, err1 := ToEng(exceptionRow[1])
				if err1 != nil {
					log.Println("From checkForExceptions (newWord)----")
					log.Println(err1.Error())
					panic(err1)
				}
				//word = strings.Replace(word, exception, newWord, 1)
				// this might be the holy grail of solutions
				//return word, true
				return newWord, true
			}
		} else if i >= lineStart-1 && val[0] == lineEnd {
			return word, false
		}
	}
	// check if the first word exists in word
	// if it does then transliterate the second word then replace
	// return replaced
	// no suffix removal or last word conversion rule check here
	return word, false
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
