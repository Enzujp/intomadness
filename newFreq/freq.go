package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)
	for s.Scan(){
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[w]++
		}
	if err:= s.Err(); err != nil {
		fmt.Printf("Error: %v", err)
	}
	}
	return freqs, nil
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0{
		return "", fmt.Errorf("error: empty map")
	}

	maxW, maxN := "", 0
	for word, count := range freqs{
		if count > maxN{
			maxW, maxN = word, count
		}
	}
	return maxW, nil
}

func commonWord(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	return maxWord(freqs)
}

func main(){
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	defer file.Close()
	word, err := commonWord(file)
	if err != nil{
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("The most common word is : ", word)
}