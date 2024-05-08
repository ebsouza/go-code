package main

import (
	"fmt"
	"testing"
)

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(i)
	}
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
}

func BenchmarkPrimeNumbersWithVariableInput(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				primeNumbers(v.input)
			}
		})
	}
}
