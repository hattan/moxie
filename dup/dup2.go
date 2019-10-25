package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup2() {
	counts := make(map[string]map[string]int)
	files := os.Args[2:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for _, foo := range counts {
		for line, n := range foo {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}
		if counts[fileName] == nil {
			counts[fileName] = make(map[string]int)
		}

		counts[fileName][line]++
	}
}
