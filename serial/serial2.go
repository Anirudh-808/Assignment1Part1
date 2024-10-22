package main

import (
	"fmt"

	"os"

	"strings"

	//used the bufio module to scan line by line
	"bufio"

	"log"
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
	vowels := "aeiou"
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	line := ""
	for scanner.Scan() {
		line = scanner.Text() // Get the current line as a string
		for _, char := range line {
			if strings.ContainsRune(vowels, strings.ToLower(char)) {
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Handle any errors that occurred during scanning
	}

	fmt.Println(count)
}
