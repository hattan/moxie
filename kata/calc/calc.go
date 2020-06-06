package calc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Split(r rune) bool {
	return r == ',' || r == 10
}

var callCount = 0

func Reset() {
	callCount = 0
}

func Sum(input string) (int, error) {
	callCount++

	if input == "" {
		return 0, nil
	}
	var tokens []string

	if strings.HasPrefix(input, "//") {
		tokens = strings.Split(input[4:], input[2:3])
	} else if strings.HasPrefix(input, "//[") {
		re := regexp.MustCompile("\\[(.*?)\\]")
		match := re.FindStringSubmatch(input)
		fmt.Printf("input=%s match=%s\n\n", input, match)
		tokens = strings.Split(input[4:], match[1])
	} else {
		tokens = strings.FieldsFunc(input, Split)
	}

	var sb strings.Builder
	total := 0
	for _, token := range tokens {
		i, _ := strconv.Atoi(token)
		if i < 0 {
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(i))
		}
		if i < 1000 {
			total = total + i
		}
	}
	negatives := sb.String()
	if len(negatives) > 0 {
		return 0, errors.New("negatives not allowed " + negatives)
	}
	return total, nil
}

func GetCalledCount() int {
	return callCount

}
