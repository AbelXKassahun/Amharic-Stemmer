package main

import (
	"fmt"
	"github.com/AbelXKassahun/Amharic-Stemmer/stemmer"
)

func main() {
	//demo := []string{
	//	"አንበሶቹ",
	//	"አንበሳ",
	//	"አበባቸው",
	//	"አበቦቻቸው",
	//	"ሰማዕታት",
	//	"ጎበኛላችሁ",
	//	"ትሄዳላችሁ",
	//	"ታመሻላችሁ",
	//	"ሞታላችሁ",
	//	"ሌቦች",
	//	"ዝንጀሮቻችን",
	//	"ደበደባቸው",
	//	"መፅሀፍት",
	//	"ፈላለገ",
	//	"ቆራረጠ",
	//	"ተመቻቸ",
	//	"ሸዋወዳቸው",
	//	"ቅጠላቅጠል",
	//	"ቅጠላቅጠሎች",
	//	"ለበሰ",
	//	"ለመለወጥ",
	//	"ለመብራት",
	//	"ለነገ",
	//	"ለቀይ",
	//	"ለእሳት",
	//	"ትምህርት",
	//	"ተነሳን",
	//	"ሄድን",
	//	"በላን",
	//	"ጠጣን",
	//	"ከፈልን",
	//	"ዐልበላም",
	//	"ከስራ",
	//	"ከጥናጥ",
	//	"ከደሞዝ",
	//	"ከበሮ",
	//	"ከነቤተሰቦቻቸው",
	//	"እባብ",
	//	"ብረታብረት",
	//	"ቀጫጭን",
	//	"የተባበሩት",
	//	"እንደቤተሰቦቻችን",
	//}
	//
	//for _, val := range demo {
	//	arr, _ := stemmer.Stem(val)
	//	fmt.Println(val, arr) //ታመሻላችሁ [ኣመሸ ኣምሽ] መሸ ምሽ ቀጫጭን [ቀጫጨ ቅጭጭ]
	//}

	arr, _ := stemmer.Stem("እንደቤተሰቦቻችን") // ሐ:ሙ፤ስ።  ሐ:*/ሙ፤ስ። [ሰማእት ስምእት]
	//arr := stemmer.KaHandler("kafaln")
	fmt.Println(arr)

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
// 		normedInput, ok := utils.inputCheck(word, false)
// 		if ok != nil {
// 			for _, val := range ok {
// 				fmt.Println(val)
// 			}
// 			// return
// 			panic("invalid input")
// 		}
// 		word = normedInput
// 	} else if fileName != "" {
// 		val, ok := utils.inputCheck(word, true)
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
