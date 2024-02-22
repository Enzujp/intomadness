package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// getting the regular expression for what is required, to have access to the functions

// "Who's on first ?" -> "Who s on first "
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)


func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)

	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line, s.Text() gives us one word from the page "-1" means to find all occurences of the pattern
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0{
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs{
		if count > maxN {
			maxN, maxW = count, word
		}
	}	
	return maxW, nil
	 
} 

// io.reader is a way to read a book. Bufio.NewScanner is like a magnifying glass that lets us read line after line, word after word
func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

func main(){
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println("The most common word in this book is :", w)

}
