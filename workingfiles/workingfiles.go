package main

import (
	"io"
	"os"
)

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil{
		return "", err
	}

	
}
//Rest in the Lord