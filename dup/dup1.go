package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count:%d\tline:%s\n", n, line)
		}
	}
}
