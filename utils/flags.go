package utils

import(
	"flag"
	// "fmt"
)

func InitializeFlags() (string, string){
	const (
		usageWord  = "Input a single word to be stemmed (dont use spaces)"
		usageWords = "Input filename that contains words to be stemmed"
	)

	var word string
	var fileName string

	flag.StringVar(&word, "w", "", usageWord)
	flag.StringVar(&word, "word", "", usageWord)
	flag.StringVar(&fileName, "f", "", usageWords)
	flag.StringVar(&fileName, "file", "", usageWords)

	flag.Parse()

	return word, fileName
}