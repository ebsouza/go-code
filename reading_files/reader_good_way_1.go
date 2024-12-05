package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			fmt.Printf("ReadLine: %q\n", line)
		}
		if err != nil {
			break
		}
	}
}
