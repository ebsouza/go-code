package main

import (
	"log"
	"os"
	"fmt"
	"bytes"
)

func main() {
	file, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	word := []byte{}
	breakLine := []byte("\n")

	for _, data := range file {
		if !bytes.Equal([]byte{data}, breakLine) {
			word = append(word, data)
		} else {
			fmt.Printf("ReadLine: %q\n", word)
			word = word[:0]
		}
	}
}