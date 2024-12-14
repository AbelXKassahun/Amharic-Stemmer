package stemmer

import (
	"fmt"
	"strings"
	// "sync"
	"github.com/AbelXKassahun/Amharic-Stemmer/utils"
)

func ToAmh(word string) (string, error) {
	var amhWord string
	arr := EngToArray(word)

	for _, val := range arr {
		if result, ok := SingleCharEngToAmh(val); ok != nil {
			fmt.Printf("\n")
			return "", ok
		} else {
			amhWord += result
		}
	}
	return amhWord, nil
}

func EngToArray(word string) []string {
	vowels := []string{"a", "u", "i", "ā", "é", "e", "o"}
	tokens := strings.Split(word, "")

	var arr = make([]string, len(tokens))
	var isVowel bool
	var prevWordIsCons bool

	for i, token := range tokens {

		for _, vowel := range vowels {
			if token == vowel {
				isVowel = true
				break
			}
			isVowel = false
		}

		// check if the token is a vowel
		if isVowel {
			// if the previous token is not a consonant,
			// then this vowel is either the first token or its predecessor is also a vowel
			if !prevWordIsCons { // አ or አእ
				arr[i] = token
			} else { // if the previous token is a consonant then this vowel is its pair and should be concatenated
				arr[i-1] += token
			}
			prevWordIsCons = false
		} else if !isVowel {
			arr[i] = token
			prevWordIsCons = true
		}
	}

	// remove spaces
	var sanitized string
	for i, val := range arr {
		if val != "" {
			if i != len(arr)-1 {
				sanitized += val + "|"
			} else {
				sanitized += val
			}
		}
	}
	// remove trailing "|"
	sanitized = strings.TrimRight(sanitized, "|")

	return strings.Split(sanitized, "|")
}

func ToEng(word string) (string, error) {
	var engWord string
	letters := strings.Split(word, "")

	for _, val := range letters {
		result, ok := SingleCharAmhToEng(val)
		if ok != nil {
			fmt.Printf("\n")
			return "", ok
		} else {
			engWord += result
			// fmt.Println(result)
		}
	}

	return engWord, nil
}

func SingleCharAmhToEng(letter string) (string, error) {
	var result string
	found := false
	for _, val := range *utils.ReturnLetters() {
		for _, val2 := range val {
			words := strings.Split(strings.TrimLeft(val2, " "), " ")
			if words[0] == letter {
				found = true
				result = words[1]
			}
		}
	}

	// if result != "" {
	if found {
		return result, nil
	} else {
		return "", fmt.Errorf("%v is not found (%v is not an amharic letter)", letter, letter)
	}
}

func SingleCharEngToAmh(letter string) (string, error) {
	var result string
	for _, val := range *utils.ReturnLetters() {
		for _, val2 := range val {
			// var words = make([]string, 3)
			words := strings.Split(strings.TrimLeft(val2, " "), " ")
			if strings.ContainsAny(words[1], "/") {
				duoTrans := strings.Split(words[1], "/")
				words = append(words, duoTrans...)
				// words[1], words[2] = duoTrans[0], duoTrans[1]

				if words[1] == letter || words[2] == letter {
					result = words[0]
				}
			} else {
				if words[1] == letter {
					result = words[0]
				}
			}
		}
	}

	if result != "" {
		return result, nil
	} else {
		return "", fmt.Errorf("%v - not found (no translation for %v)", letter, letter)
	}
}

func removeVowels(word string) string {
	vowels := []string{"a", "u", "i", "ā", "é", "e", "o"}
	tokens := strings.Split(word, "")

	var arr = make([]string, len(tokens))
	var isVowel bool
	var prevWordIsCons bool

	for i, token := range tokens {

		for _, vowel := range vowels {
			if token == vowel {
				isVowel = true
				break
			}
			isVowel = false
		}

		// check if the token is a vowel
		if isVowel {
			// if the previous token is not a consonant,
			// then this vowel is either the first token or its predecessor is also a vowel
			if !prevWordIsCons { // አ or አእ
				arr[i] = token
			}
			// if the previous token is a consonant then this vowel is its pair and should be abandoned
			prevWordIsCons = false
		} else if !isVowel {
			arr[i] = token
			prevWordIsCons = true
		}
	}

	return strings.ReplaceAll(strings.Join(arr, ""), " ", "")
}

// this function spawns 34 goroutines which is the number of rows in the letters.csv file
// it was meant to be faster than brute force comparison
// but it's not, the average execution time is 500 microseconds
// while the brute force function averages around 180 microseconds
// this is a stupid approach even if it was faster than brute force comparison

// func LookForLetter2(record [][]string, letter string) (string, error){
// 	var wg sync.WaitGroup
// 	end := make(chan int)
// 	var result string
// 	found := false
// 	// mu := &sync.Mutex{}

// 	for _, val := range record {
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			for _, val2 := range val {
// 				select{
// 				case <- end:
// 					// fmt.Println("closed")
// 					return
// 				default:
// 					words := strings.Split(strings.TrimLeft(val2, " "), " ")
// 					if words[0] == letter{
// 						// mu.Lock()
// 						result = words[1]
// 						found = true
// 						// mu.Unlock()
// 						close(end)
// 						return
// 						// return words[1], nil
// 					}
// 				}
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	if found {
// 		return result, nil
// 	}else{
// 		return "", fmt.Errorf("%v - not found (might not be an amharic letter)", letter)
// 	}
// }
