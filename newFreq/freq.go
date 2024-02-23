package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// counting words in the book and returning the word with the highest frequency
// Glory be to God the father almighty!

// Using sherlock.txt as reference

// write readfile function

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func findFrequency(r io.Reader) (map[string]int, error) {
	// get scanner and pass it io reader
	s := bufio.NewScanner(r)
	// create map to store word and count from read file
	frequency := make(map[string]int) // word -> count

	// since the file is going to be passed into the reader we, can scan through using buffio
	for s.Scan(){
		words := wordRe.FindAllString(s.Text(), -1) // find all strings matching regExp pattern 
		for _, word := range words {
			frequency[word]++
		}
		// handle any errors that may occur during reading
		if err := s.Err(); err != nil{
			return nil, err
		}
	}
	return frequency, nil
}

func maxWord(frequency map[string]int) (string, error) {
	if len(frequency) == 0 {
		return "", fmt.Errorf("error: nothing contained in Map")
	}

	maxW, maxN := "", 0

	for word, count:= range frequency {
		if count > maxN {
			maxW, maxN = word, count 
		}

	}
	return maxW, nil
}

func mostCommonWord(r io.Reader) (string, error) {
	freqs, err := findFrequency(r)
	if err != nil {
		return "", fmt.Errorf("could not process due to error: %v", err)
	}

	return maxWord(freqs)
}
func main(){
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Printf("error encountered in opening file: %v", err)
	}
	defer file.Close()

	word, err := mostCommonWord(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("The most common word in the book is : %v\n", word)

}