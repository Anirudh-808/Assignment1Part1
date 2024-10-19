package main

import (
	"fmt"

	"os"
	//importing the bufio module(used ChatGPT for this)
	//used for reading the file
	"bufio"
	//importing the strings module(used ChatGPT for this)
	//used to modify strings
	"strings"
)

// making a func to check if a char is vowel or not
// func takes input string and return bool
func isVowel(r string) bool {
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

func main() {
	//the below is used to import the file big.txt
	//Using Open and nor Readfile due to the bufio module
	dat, _ := os.Open("./big.txt")

	//set a counter
	count := 0

	//the Newscanner function creates a new scanner to scan the doc
	//it defaults to reading line by line
	scanner := bufio.NewScanner(dat)
	//iterate over the data in file and check
	for scanner.Scan() {
		line := scanner.Text()
		for _, l := range line {
			if isVowel(string(l)) {
				count++
			}
		}
	}

	fmt.Println(count)

}
