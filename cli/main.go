package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {

	byLine := flag.Bool("l", false, "Count lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *byLine))
}
