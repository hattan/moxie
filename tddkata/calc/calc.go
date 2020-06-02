package calc

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	sum := 0
	data := strings.Split(numbers, ",")

	for _, number := range data {
		num, _ := strconv.Atoi(number)
		sum = sum + num
	}
	return sum
}
