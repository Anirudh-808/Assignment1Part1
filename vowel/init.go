package vowel

import (
	//importing the strings module(used ChatGPT for this)
	//used to modify strings
	"strings"
)

func IsVowel(r string) bool {
	L := []string{"a", "e", "i", "o", "u"}
	//use the Tolower method to lower the input string
	r_new := strings.ToLower(r)

	for _, v := range L {
		if v == r_new {
			return true
		}
	}

	return false
}

var Count int = 0

func CheckifVowelParallel(a string) int {
	for _, l := range a {
		if IsVowel(string(l)) {
			Count++
		}
	}

	return Count

}

func CheckifVowelSerial(a string) int {
	for _, l := range a {
		if IsVowel(string(l)) {
			Count++
		}
	}

	return Count

}
