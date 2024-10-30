package main

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
)
func main() {
	smtn := stemmer.Stem("liju man endehone tawkalachu? esu ")
	fmt.Println(smtn)
}