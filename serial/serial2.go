package main

import (
	"fmt"

	"os"

	"strings"

	//used the bufio module to scan line by line
	"bufio"
)

// func CheckIfVowel(s string) int {
// 	L := []string{"a", "e", "i", "o", "u"}
// 	counter := 0
// 	for _, l := range L {
// 		if l == strings.ToLower(s) {
// 			counter++
// 		}
// 	}

// 	return counter
// }

func main() {
	dat, _ := os.Open("./big.txt")
	count := 0
	L := []string{"a", "e", "i", "o", "u"}

	scanner := bufio.NewScanner(dat)
	line := ""
	for scanner.Scan() {
		line = scanner.Text() // Get the current line as a string
		for _, a := range line {
			for _, l := range L {
				if l == strings.ToLower(string(a)) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
