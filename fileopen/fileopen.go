package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := openFile("http.log.gz")
	if err != nil {
		log.Fatalf("Couldnt read file : %s", err)
	}
	fmt.Println(sig)
}

func openFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
	}

	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(r)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
