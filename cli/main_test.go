package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")

	expected := 3

	result := count(b, false)

	if result != expected {
		t.Errorf("Expected: %d, Result: %d. \n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("line1 \n line2 \n line3 \n line4")

	expected := 4

	result := count(b, true)

	if result != expected {
		t.Errorf("Expected: %d, Result: %d. \n", expected, result)
	}
}
