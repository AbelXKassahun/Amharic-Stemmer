package main

import (
	"fmt"
	stemmer "github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
	// utils "github.com/AbelXKassahun/Amharic-Stemmer/utils"
	"time"
)

func main() {
	// start the timer for execution
	start := time.Now()
	
	// initialize the flags
	// word, fileName := utils.InitializeFlags()



	// else if fileName != "" {
	// 	val, ok := utils.InputCheck(word, true)
	// 	if ok != nil {
	// 		panic(ok)
	// 	}
	// 	fileName = val
	// }
	// fmt.Println(fileName)
	
	// open the amharic letters conversion table file

	/*
	መፅሃፋችን
መፅሃፋቸው
መፅሃፎች
መፅሃፎቹን
	*/

	fmt.Println(stemmer.ToEng("መፅሃፋችን"))
	fmt.Println(stemmer.ToEng("መፅሃፋቸው"))
	fmt.Println(stemmer.ToEng("መፅሃፎች"))
	fmt.Println(stemmer.ToEng("መፅሃፎቹን"))


	// calculate and displat the time it took to 
	duration := time.Since(start)
	fmt.Printf("\n\nThis code took %v", duration)
}