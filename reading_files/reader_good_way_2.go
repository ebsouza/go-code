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

	s := bufio.NewScanner(file)

	for s.Scan() {
		fmt.Println(s.Text())
	}
}
