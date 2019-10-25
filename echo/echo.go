package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func getStringJoin(args []string) string {
	return strings.Join(os.Args[1:], " ")
}

func getStringLoop(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func getStringLoop2(args []string) string {
	s := ""
	for i := 1; i < len(args); i++ {
		s += " " + args[i]
	}
	return s
}

func getProgramName() string {
	index := strings.LastIndex(os.Args[0], "/")
	me := os.Args[0]
	return me[index+1 : len(me)]
}

func main() {
	debugPtr := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	me := getProgramName()
	text := getStringJoin(os.Args)

	if !(*debugPtr) {
		fmt.Printf("%s\n", text)
	} else {
		fmt.Printf("%s %s\n", me, os.Args[1])
		fmt.Printf("%s\n", text)
	}
}
