package calc

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	message       string
	testId        string
	input         string
	expected      int
	expectedError error
}{
	{
		message:       "Empty string should return 0.",
		testId:        "1",
		input:         "",
		expected:      0,
		expectedError: nil,
	},
	{
		message:       "Single character should return correct digit.",
		testId:        "2",
		input:         "5",
		expected:      5,
		expectedError: nil,
	},
	{
		message:       "Two Characters returns sum of both.",
		testId:        "3",
		input:         "1,2",
		expected:      3,
		expectedError: nil,
	},
	{
		message:       "Three Characters returns sum of all.",
		testId:        "4",
		input:         "1,5,10",
		expected:      16,
		expectedError: nil,
	},
	{
		message:       "Supports newline as a delimeter",
		testId:        "5",
		input:         "1\n2,3",
		expected:      6,
		expectedError: nil,
	},
	{
		message:       "Support different delimiters",
		testId:        "6",
		input:         "//;\n1;2",
		expected:      3,
		expectedError: nil,
	},
	{
		message:       "Negative Numbers returns error",
		testId:        "7",
		input:         "-1,5",
		expected:      3,
		expectedError: errors.New("negative numbers not allowed found:(-1)"),
	},
	{
		message:       "Multiple Negative Numbers appear in error",
		testId:        "7",
		input:         "-1,5,-20,-2,4",
		expected:      3,
		expectedError: errors.New("negative numbers not allowed found:(-1,-20,-2)"),
	},
}

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func TestAdd(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.testId, func(t *testing.T) {
			//act
			result, err := Add(test.input)

			//assert
			switch {
			case test.expectedError != nil:

				if assert.Error(t, err) {
					assert.Equal(t, test.expectedError, err)
				}

			default:
				assert.Equal(t, test.expected, result, test.message)
			}

		})
	}
}
