package main

import (
	"fmt"

	"os"
	//importing the bufio module(used ChatGPT for this)
	//used for reading the file
	"bufio"

	"sync"

	//importing the newly made vowel package
	"part1/vowel"
)

// making a func to check if a char is vowel or not
// func takes input string and return bool
var wg sync.WaitGroup

func main() {
	//the below is used to import the file big.txt
	//Using Open and nor Readfile due to the bufio module
	dat, _ := os.Open("./big.txt")

	//set a counter
	count := 0

	//the Newscanner function creates a new scanner to scan the doc
	//it defaults to reading line by line
	c := make(chan string)
	scanner := bufio.NewScanner(dat)
	//iterate over the data in file and check
	for scanner.Scan() {
		c <- scanner.Text()
		for line := range c {
			wg.Add(1)
			go vowel.CheckifVowelParallel(line)
			wg.Done()
		}

	}

	wg.Wait()

	fmt.Println(count)

}
