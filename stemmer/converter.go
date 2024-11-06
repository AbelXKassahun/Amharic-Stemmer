package stemmer

import(
	"fmt"
	"strings"
	// "sync"
)

func LookForLetter(record [][]string, letter string) (string, error){
	var result string
	found := false
	for _,val := range record {
		for _, val2 := range val{
			words := strings.Split(strings.TrimLeft(val2, " "), " ")
			if words[0] == letter{
				found = true
				result = words[1]
			}
		}
	}
	if found {
		return result, nil
	}else{
		return "", fmt.Errorf("%v - not found (%v is not an amharic letter)", letter, letter)
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