//Base Conversion package - This is probably over engineered :)
package allyourbase

import (
	"errors"
)

// Implements exponents. returns a^b
func pow(a int, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	return result
}

// Reverses the order of elements in an Array
func reverse(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

// Returns a number representing the input in base10 (decimal)
func getDecimal(digits []int, inputBase int) (int, error) {
	decimalTotal := 0
	for i, digit := range digits {
		if digit < 0 || digit >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}

		position := ((len(digits) - 1) - i)
		decimalTotal = decimalTotal + (digit * pow(inputBase, position))
	}
	return decimalTotal, nil
}

//Converts a base10 number to the desired output base
func ConvertToBaseFromDecimal(decimal int, outputBase int) []int {
	results := make([]int, 0)
	for decimal > 0 {
		a := decimal % outputBase
		decimal = decimal / outputBase
		results = append(results, a)
	}

	if len(results) == 0 {
		results = []int{0}
	}
	return results
}

//Converts a given array of integers from one base to another.
func ConvertToBase(inputBase int, digits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	decimalTotal, err := getDecimal(digits, inputBase)
	if err != nil {
		return nil, err
	}

	results := ConvertToBaseFromDecimal(decimalTotal, outputBase)
	if len(results) == 0 {
		results = []int{0}
	}

	return reverse(results), nil
}
