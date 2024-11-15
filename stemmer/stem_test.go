package stemmer

import (
	// "fmt"
	"reflect"
	"testing"

	// "golang.org/x/tools/go/expect"
)

func TestConvHouse(t *testing.T) {
	testCases := []struct {
		letter   string
		house    int
		expected string
	}{
		{"ha", 6, "he"},
		{"wa", 3, "wi"},
		{"si", 1, "sa"},
		{"ze", 3, "zi"},
	}

	for _, tc := range testCases {
		result := convertToDiffHouse(tc.letter, tc.house)
		if result != tc.expected {
			t.Errorf("FAIL - convertToDiffHouse(%v, %d) = %v; expected %v", tc.letter, tc.house, result, tc.expected)
		} else {
			t.Logf("PASS - convertTodifffHouse(%v, %d) = %v; expected %v", tc.letter, tc.house, result, tc.expected)
		}
	}
}

func TestSample(t *testing.T) {
	testCases := []struct{
		input string
		expected []string
	}{
		{"abel", []string{"abel"}},
		{"kass;ahun", []string{"kass", "ahun"}},
	}

	for _, tc := range testCases {
		result := sample(tc.input)
		if reflect.DeepEqual(result, tc.expected) {
			t.Logf("PASS - sample(%v) = %v; expected %v", tc.input, result, tc.expected)
		} else {
			t.Errorf("FAIL - sample(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestRemoveShortSuffix(t *testing.T) {
	testCases := []struct{
		bareSuffix []string
		word []string
		rules []string
		expected []string
	}{
		{[]string{"we"}, []string{"wa","ré", "we"}, []string{"no-split"}, []string{"wa","ré"}},
		{[]string{"u"}, []string{"ho", "té", "lu"}, []string{"split", "6"}, []string{"ho", "té", "le"}},
		{[]string{"u"}, []string{"ma", "ṡe", "ha", "fu"}, []string{"split", "6"}, []string{"ma", "ṡe", "ha", "fe"}},
	}

	for _, tc := range testCases {
		result := removeShortSuffix(tc.bareSuffix, tc.word, tc.rules)
		if reflect.DeepEqual(result, tc.expected){
			t.Logf("PASS - removeShortSuffix(%v, %v ,%v) = %v; expected %v", tc.bareSuffix, tc.word, tc.rules, result, tc.expected)
		} else {
			t.Errorf("FAIL - removeShortSuffix(%v, %v ,%v) = %v; expected %v", tc.bareSuffix, tc.word, tc.rules, result, tc.expected)
		}
	}
}

func TestFindSuffixSubstringIndex(t *testing.T) {
	testCases := []struct {
		bareSuffix []string
		word []string
		expected []int
	}{
		{[]string{"ā","če","ne"}, []string{"ma", "ṡe", "hā", "fā", "če", "ne"}, []int{3, 5}},
		{[]string{"ā","ča","we"}, []string{"ma", "ṡe", "hā", "fā", "ča", "we"}, []int{3, 5}},
		{[]string{"o","če"}, []string{"ma", "ṡe", "hā", "fo", "če"}, []int{3, 4}},
		// {[]string{"o","če","ne"}, []string{"ma", "ṡe", "ha", "fu"}, []string{"ma", "ṡe", "ha", "fe"}},
		{[]string{"o","ču","ne"}, []string{"ma", "ṡe", "hā", "fo", "ču", "ne"}, []int{3, 5}},
	}
	// "ma", "ṡe", "hā", "fe"
	for _, tc := range testCases {
		start, end := findSuffixSubstringIndex(tc.bareSuffix, tc.word)
		var result []int
		result = append(result, start, end)
		if reflect.DeepEqual(result, tc.expected){
			t.Logf("PASS - findSuffixSubstringIndex(%v, %v) = start : %v, end : %v; expected %v", tc.bareSuffix, tc.word, start, end, tc.expected)
		} else {
			t.Errorf("FAIL - findSuffixSubstringIndex(%v, %v) = start : %v, end : %v; expected %v", tc.bareSuffix, tc.word, start, end, tc.expected)
		}
	}
}

func TestRemoveLongSuffix(t *testing.T){
	testCases := []struct{
		bareSuffix []string
		word []string
		rules []string
		expected []string
	}{
		{[]string{"ā","če","ne"}, []string{"ma", "ṡe", "hā", "fā", "če", "ne"}, []string{"split", "6"}, []string{"ma", "ṡe", "hā", "fe"}},
		{[]string{"ā","ča","we"}, []string{"ma", "ṡe", "hā", "fā", "ča", "we"}, []string{"split", "6"}, []string{"ma", "ṡe", "hā", "fe"}},
		{[]string{"o","če"}, []string{"ma", "ṡe", "hā", "fo", "če"}, []string{"split", "6"}, []string{"ma", "ṡe", "hā", "fe"}},
		{[]string{"o","ču","ne"}, []string{"ma", "ṡe", "hā", "fo", "ču", "ne"},[]string{"split", "6"}, []string{"ma", "ṡe", "hā", "fe"}},
	}

	for _, tc := range testCases {
		result := removeLongSuffix(tc.bareSuffix, tc.word, tc.rules)
		// fmt.Printf("%v, %v", result, tc.expected)
		if reflect.DeepEqual(result, tc.expected){
			t.Logf("PASS - removeLongSuffix(%v, %v ,%v) = %v; expected %v", tc.bareSuffix, tc.word, tc.rules, result, tc.expected)
		} else {
			t.Errorf("FAIL - removeLongSuffix(%v, %v ,%v) = %v; expected %v", tc.bareSuffix, tc.word, tc.rules, result, tc.expected)
		}
	}
}

func TestStem(t *testing.T) {
	testCases := []struct{
		word string
		expected string
	}{
		{"መፅሃፋችን", "መፅሃፍ"},
	}

	for _, tc := range testCases{
		result := Stem(tc.word)
		if result == tc.expected {
			t.Logf("PASS - Stem(%v) result: %v, expected: %v", tc.word, result, tc.expected)
		} else {
			t.Errorf("FAIL - Stem(%v) result: %v, expected: %v", tc.word, result, tc.expected)
		}
	}
}