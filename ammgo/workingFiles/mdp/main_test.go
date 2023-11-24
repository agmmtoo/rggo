package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := parseContent(input)

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}
