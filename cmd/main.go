package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Frequency map[rune]int

// NewFrequency returns a new rune frequency map for the given file
func NewFrequency(file *os.File) Frequency {
	freq := Frequency{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			freq.increment(char)
		}
	}

	return freq
}

// increment will increment frequency of a unicode in the map
// don't need pointer receiver as map functions already apply result by reference
func (f Frequency) increment(char rune) {
	if _, ok := f[char]; ok {
		f[char]++
	} else {
		f[char] = 1
	}
}

func main(){
	path := flag.String("f", "", "file path of text to encode/decode")
	flag.Parse()

	if *path == "" {
		fmt.Println("Error: provide file path")
		os.Exit(1)
	}
	fmt.Printf("File to process: %s\n", *path)

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	freq := NewFrequency(file)

	for k, v := range freq {
		fmt.Printf("%s => %v\n", string(k), v)
	}
}
