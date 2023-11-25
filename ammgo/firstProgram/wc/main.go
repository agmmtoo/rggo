package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type filenamesValue []string

func (v *filenamesValue) String() string {
	return fmt.Sprintf("files: %s", strings.Join(*v, ","))
}

func (v *filenamesValue) Set(value string) error {
	for _, filename := range strings.Split(value, ",") {
		*v = append(*v, filename)
	}
	return nil
}

func main() {
	var files filenamesValue
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Var(&files, "files", "List of comma separated file names")
	flag.Parse()

	if len(files) > 0 {
		var totalCount int
		for _, fileName := range files {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			totalCount += count(file, *lines, *bytes)
			file.Close()
		}
		fmt.Println(totalCount)
	} else {
		fmt.Println(count(os.Stdin, *lines, *bytes))
	}
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
