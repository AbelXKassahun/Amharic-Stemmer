package stemmer

import (
	"fmt"
	"strings"
	// "sync"
	utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
)

func ToAmh(word []string) (string, error) {
	var amhWord string

	for _, val := range word {
		if result, ok := SingleCharEngToAmh(val); ok != nil {
			fmt.Printf("\n")
			return "", ok
		} else {
			amhWord += result
			fmt.Println(result)
		}
	}
	return amhWord, nil
}

func ToEng(word string) ([]string, error) {
	var engWord []string
	letters := strings.Split(word, "")

	for _, val := range letters {
		if result, ok := SingleCharAmhToEng(val); ok != nil {
			fmt.Printf("\n")
			return []string{}, ok
		} else {
			engWord = append(engWord, result)
			fmt.Println(result)
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

// this function spawns 34 goroutines which is the number of rows in the letters.csv file
// it was meant to be faster than brute force comparison
// but its not, the average execution time is 500 microseconds
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
