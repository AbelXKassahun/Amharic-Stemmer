package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func InputCheck(word string, isFileName bool) (string, []error){
	var err []error

	words := strings.Split(strings.TrimSpace(word), " ")
	
	if words[0] == "" {
		err = append(err, fmt.Errorf("no input detected"))
		return "", err
	} else if isFileName {
		return strings.ReplaceAll(word, " ", ""), nil
	}
	
	letters := strings.Split(words[0], "")
	
	// pattern := `^(?=.*[A-Za-z])(?=.*\d)(?=.*[^\w\s]).+$`
	alphabets := `[A-Za-z]`
	digits := `[0-9]`
	symbols := `[^\p{L}\s]`

	isLetter := regexp.MustCompile(alphabets)
	isDigit := regexp.MustCompile(digits)
	isSymbol := regexp.MustCompile(symbols)

	for _,val := range letters {
		if isLetter.MatchString(val) {
			err = append(err, fmt.Errorf("%s - is not an amharic letter, cant use english alphabets", val))
		} else if isDigit.MatchString(val){
			err = append(err, fmt.Errorf("%s - is not an amharic letter, cant use digits", val))
		} else if isSymbol.MatchString(val){
			err = append(err, fmt.Errorf("%s - is not an amharic letter, cant use symbols", val))
		}
	}

	if len(err) > 0 {
		return "", err
	}

	return word, nil
}