package stemmer

import (
	"fmt"
	"regexp"
)

func RemoveInfix(word string) string {
	if infix1removal, found := infixPattern1Removal(word); found {
		return infix1removal
	} else if infix2removal, found := infixPattern2Removal(word); found {
		return infix2removal
	}
	return word
}

func infixPattern1Removal(word string) (string, bool) {
	pattern := `([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([a])([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([ā])([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([a])`
	re := regexp.MustCompile(pattern)
	var found bool
	if re.MatchString(word) {
		found = true
		match := re.FindString(word)
		fmt.Println("matches ", match)
	}
	// Replace the matched pattern by removing CF (3rd and 4th groups).
	result := re.ReplaceAllStringFunc(word, func(match string) string {
		subMatches := re.FindStringSubmatch(match)
		for _, val := range subMatches {
			fmt.Printf("%v, ", val)
		}
		if len(subMatches) >= 7 {
			// Reconstruct: keep 1st, 2nd, 5th, and 6th groups only
			fmt.Println(subMatches[1] + ", " + subMatches[2] + ", " + subMatches[5] + ", " + subMatches[6])
			return subMatches[1] + subMatches[2] + subMatches[5] + subMatches[6]
		}
		return match
	})

	return result, found
}

func infixPattern2Removal(word string) (string, bool) {
	pattern := `([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([a])([hlḥmśrsšqbtčnñkxwzžydǧgṭċp̣ṣṡfpv])([ā])`
	re := regexp.MustCompile(pattern)
	var found bool
	if re.MatchString(word) {
		found = true
		match := re.FindString(word)
		fmt.Println("matches ", match)
	}
	// Replace the matched pattern by removing CF (3rd and 4th groups).
	result := re.ReplaceAllStringFunc(word, func(match string) string {
		subMatches := re.FindStringSubmatch(match)

		if len(subMatches) >= 6 {
			return ""
		}
		return match
	})

	return result, found
}
