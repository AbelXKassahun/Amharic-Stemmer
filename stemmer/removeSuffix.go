package stemmer

import (
	"strconv"
	"strings"

	utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
)

// this works for chained suffixes
func RemoveSuffix(word []string) []string {
	var processedWord []string
	processedWord = append(processedWord, word...)
	for i, val := range *utils.ReturnSuffixList() {
		if i != 0 {
			bareSuffix := strings.Split(val[0], "|")
			rules := strings.Split(strings.TrimSpace(val[1]), "|")
			// splittedSuffix := strings.Split(strings.ReplaceAll(string(val2[0]), "|",""), "")
			if len(bareSuffix) > 1  && len(processedWord) >= 3 { // 1
				processedWord = removeLongSuffix(bareSuffix, processedWord, rules)
			} else {
				if len(processedWord) > 2 {
					processedWord = removeShortSuffix(bareSuffix, processedWord, rules)
				}
			}
		}
	}
	return processedWord
}

func removeLongSuffix(bareSuffix []string, word []string, rules []string) []string{
	start, end := findSuffixSubstringIndex(bareSuffix, word)
	var temp []string
	if start != -1 && end !=-1 {
		// split the letter at the start index and then replace it
		// according to the rules 
		temp = append(temp, word[:start]...)
		if rules[0] == "split" {
			// copy(temp, word[:start])
			house, _ := strconv.Atoi(rules[1])
			converted := convertToDiffHouse(word[start], house)
			
			temp = append(temp, converted)
			temp = append(temp, word[end+1:]...)
			return temp
		} else {
			temp = append(temp, word[end+1:]...)
			return temp
		}
	}

	return word
}

// substring = contingous
func findSuffixSubstringIndex(bareSuffix []string, word []string) (int, int) {
	i, j := 1, 0
	start, end := -1, -1
	for i < len(word) {
		if j == 0 {
			a := strings.Split(word[i], "")
			for _, val := range a {
				if val == bareSuffix[j] {
					start = i
					i++; j++
				}
			}
			// a replacement of continue lol
			if start == -1 {
				i++
			}
		} else {
			if word[i] == bareSuffix[j] {
				// if the last suffix letter and the last letter of the word mathc
				// and the start index has already been found
				// then end is the last index of the word array 
				if j == len(bareSuffix)-1 && start != -1 {
					end = i
					return start, end
				}
				// we update both index if they both match
				i++; j++
			} else {
				// we dont increment i here
				j = 0;
				start, end = -1, -1
			}
		}
	}
	return start, end
}

func removeShortSuffix(bareSuffix []string, word []string, rules []string) []string {
	if rules[0] == "split" {
		lastWordSplit := strings.Split(word[len(word)-1], "")
		for _, val := range lastWordSplit {
			if val == bareSuffix[0] {
				var temp []string
				
				house, _ := strconv.Atoi(rules[1])
				
				temp = append(temp, word[:len(word)-1]...)
				temp = append(temp, convertToDiffHouse(word[len(word)-1], house))
				return temp
			}
		}
	} else if rules[0] == "no-split" {
		if word[len(word)-1] == bareSuffix[0] {
			// remove the suffix from the word without conversion
			return word[:len(word)-1]
		}
	}
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


// wrote this function to test out gotest
func sample(a string) []string{
	return strings.Split(a, ";")
}