package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	if numbers == "" {
		return 0, nil
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
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}

		if num < 0 {
			return 0, errors.New("negative numbers not allowed")
		}
		sum = sum + num
	}
	return sum, nil
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}
