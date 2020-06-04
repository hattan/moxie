package calc

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	delimeters := ","

	if len(numbers) > 1 && numbers[0:2] == "//" {
		index := strings.Index(numbers, "\n")
		delimeters = delimeters + numbers[2:index]
		numbers = numbers[index+1:]
	} else {
		delimeters = delimeters + "\n"
	}

	sum := 0
	data := splitAny(numbers, delimeters)
	fmt.Println(data)
	for _, number := range data {
		num, _ := strconv.Atoi(number)
		sum = sum + num
	}
	return sum
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}
