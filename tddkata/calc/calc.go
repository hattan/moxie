package calc

import (
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
	negatives := ""
	for _, number := range data {
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}

		if num < 0 {
			negatives = includeDisplayComma(negatives) + number
		}

		sum = sum + num
	}
	if negatives != "" {
		return 0, fmt.Errorf("negative numbers not allowed found:(%s)", negatives)
	}
	return sum, nil
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func includeDisplayComma(input string) string {
	if input != "" {
		input = input + ","
	}
	return input
}
