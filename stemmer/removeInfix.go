package stemmer

import (
	"regexp"
	"strings"
)

//var Consonant = `([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])`

func RemoveInfix(word string) string {
	if infixRemoval, found := infixPattern1Removal(word); found {
		return infixRemoval
	} else if infixRemoval, found := infixPattern2Removal(word); found {
		return infixRemoval
	} else if infixRemoval, found := infixExceptionRemoval(word); found {
		return infixRemoval
	}
	return word
}

func infixPattern1Removal(word string) (string, bool) {
	pattern := `([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ao])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ā])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])`
	re := regexp.MustCompile(pattern)
	var found bool
	if re.MatchString(word) {
		found = true
		//match := re.FindString(word)
		//fmt.Println("matches ", match)
	}
	// Replace the matched pattern by removing CF (3rd and 4th groups).
	result := re.ReplaceAllStringFunc(word, func(match string) string {
		subMatches := re.FindStringSubmatch(match)

		if len(subMatches) >= 8 {
			// Reconstruct: keep 1st, 2nd, 5th, and 6th groups only
			//fmt.Println(subMatches[1] + ", " + subMatches[2] + ", " + subMatches[5] + ", " + subMatches[6])
			return subMatches[1] + subMatches[2] + subMatches[5] + subMatches[6] + subMatches[7]
		}
		return match
	})

	return result, found
}

func infixPattern2Removal(word string) (string, bool) {
	// CCaCāCCaC
	pattern := `([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ā])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])([a])([ḥśščñžṭċṣṡhlmrsqbtnkxwzydfpvj9g])`

	re := regexp.MustCompile(pattern)
	var found bool
	if re.MatchString(word) {
		found = true
		//subMatches := re.FindStringSubmatch(word)
		//fmt.Println(subMatches[6] + ", " + subMatches[7] + ", " + subMatches[8] + ", " + subMatches[9])
	}
	// Replace the matched pattern by removing CF (3rd and 4th groups).
	result := re.ReplaceAllStringFunc(word, func(match string) string {
		subMatches := re.FindStringSubmatch(match)

		if len(subMatches) >= 10 {
			return subMatches[6] + subMatches[7] + subMatches[8] + subMatches[9]
		}
		return match
	})

	return result, found
}

func infixExceptionRemoval(word string) (string, bool) {
	infixExceptions := []string{
		"ተመቻ-ተመቸ",
	}
	for _, val := range infixExceptions {
		exceptionNNew := strings.Split(val, "-")
		exception, _ := ToEng(exceptionNNew[0])
		if strings.Contains(word, exception) {
			newWord, _ := ToEng(exceptionNNew[1])
			return newWord, true
		}
	}
	return word, false
}
