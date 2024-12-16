package main

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
)

func main() {
	arr, _ := stemmer.Stem("አንበሶቻችን")
	fmt.Println(arr) // [ሸወድ ሸወዳ ሸወደ ሽውድ]

	//pattern := `^(l)(a)([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpv]|ǧ|p̣)([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpv]|ǧ|p̣)([aā]?)$`
	//re := regexp.MustCompile(pattern)
	//re.ReplaceAllStringFunc("lamalaw", func(match string) string {
	//	subMatches := re.FindStringSubmatch(match)
	//	for _, val := range subMatches {
	//		fmt.Printf("%v, ", val)
	//	}
	//	fmt.Println("")
	//	return match
	//})
	//wrds := []string{"ḥ", "ś", "š", "č", "ñ", "ž", "ǧ", "ṭ", "ċ", "p̣", "ṣ", "ṡ", "h"}
	//for _, val := range wrds {
	//	re.ReplaceAllStringFunc(fmt.Sprintf("la%va%v", val, val), func(match string) string {
	//		subMatches := re.FindStringSubmatch(match)
	//		for _, val := range subMatches {
	//			fmt.Printf("%v, ", val)
	//		}
	//		fmt.Println("")
	//		return match
	//	})
	//}
}

// func main() {
// 	// start the timer for execution
// 	start := time.Now()

// 	// initialize the flags
// 	word, fileName := utils.InitializeFlags()

// 	if word == "" && fileName == "" {
// 		fmt.Println("no input detected")
// 		return
// 	} else if word != "" {
// 		normedInput, ok := utils.InputCheck(word, false)
// 		if ok != nil {
// 			for _, val := range ok {
// 				fmt.Println(val)
// 			}
// 			// return
// 			panic("invalid input")
// 		}
// 		word = normedInput
// 	} else if fileName != "" {
// 		val, ok := utils.InputCheck(word, true)
// 		if ok != nil {
// 			panic(ok)
// 		}
// 		fileName = val
// 	}
// 	fmt.Println(stemmer.Stem(word))
// 	fmt.Println(fileName)

// 	// open the amharic letters conversion table file

// 	/*
// 	መፅሃፋችን
// 	መፅሃፋቸው
// 	መፅሃፎች
// 	መፅሃፎቹን
// 	*/

// 	// fmt.Println(stemmer.ToEng("መፅሃፋችን"))
// 	// fmt.Println(stemmer.ToEng("መፅሃፋቸው"))
// 	// fmt.Println(stemmer.ToEng("መፅሃፎች"))
// 	// fmt.Println(stemmer.ToEng("መፅሃፎቹን"))

// 	// calculate and display the time it took to
// 	duration := time.Since(start)
// 	fmt.Printf("\n\nThis code took %v", duration)
// }
