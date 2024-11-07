package stemmer

import (
	"fmt"
	// "strings"
)

func Stem(record [][]string, word string) string {
	eng := ToEng(record, word)
	fmt.Println(eng)
	amh := ToAmh(record, eng)
	fmt.Println(amh)
	return amh
}

// ʼa,bé,le/l
