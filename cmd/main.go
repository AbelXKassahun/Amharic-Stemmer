package main

import (
	"fmt"
	stemmer "github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
	utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"time"
)

func main() {
	// start the timer for execution
	start := time.Now()
	
	// initialize the flags
	word, fileName := utils.InitializeFlags()

	if word == "" && fileName == "" {
		fmt.Println("no input detected")
		return
	} else if word != "" {
		normedInput, ok := utils.InputCheck(word, false)
		if ok != nil {
			for _, val := range ok {
				fmt.Println(val)
			}
			// return
			panic("invalid input") 
		}
		word = normedInput
	} 
	// else if fileName != "" {
	// 	val, ok := utils.InputCheck(word, true)
	// 	if ok != nil {
	// 		panic(ok)
	// 	}
	// 	fileName = val
	// }
	// fmt.Println(fileName)
	
	// open the amharic letters conversion table file
	record, err := utils.OpenFile("../data/letters.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stemmer.Stem(record, word))

	// calculate and displat the time it took to 
	duration := time.Since(start)
	fmt.Printf("\n\nThis code took %v", duration)
}