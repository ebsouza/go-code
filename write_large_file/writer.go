package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"math"
)

func main() {

    file, err := os.Create("words.txt")
	value := "Hello world"
    
	if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    wr := bufio.NewWriter(file)
	var counter uint64
	var limit = uint64(math.Pow(10, 2)) // exp 8 ~ 1.2Gb ; exp 9 ~ 12Gb

    for {

        wr.WriteString(value + "\n")

		if counter > limit {
			break
		}

		counter++
    }

    wr.Flush()

    fmt.Println("data written")
}