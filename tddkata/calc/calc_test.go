package calc

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	message  string
	testId   string
	input    string
	expected int
}{
	{
		message:  "Empty string should return 0.",
		testId:   "1",
		input:    "",
		expected: 0,
	},
	{
		message:  "Single character should return correct digit.",
		testId:   "2",
		input:    "5",
		expected: 5,
	},
	{
		message:  "Two Characters returns sum of both.",
		testId:   "3",
		input:    "1,2",
		expected: 3,
	},
	{
		message:  "Three Characters returns sum of all.",
		testId:   "4",
		input:    "1,5,10",
		expected: 16,
	},
}

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func TestEmptyStringReturnsZero(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.testId, func(t *testing.T) {
			//act
			result := Add(test.input)

			//assert
			assert.Equal(t, test.expected, result, test.message)
		})
	}
}
