/*
Package acronym is focused on creating Acronyms from strings
*/
package acronym

import (
	"regexp"
	"strings"
)

// Converts a phrase to its acronym.
func Abbreviate(s string) string {
	tokens := regexp.MustCompile("[\\ \\-\\_\\s]+").Split(s, -1)
	var str strings.Builder
	for _, word := range tokens {
		str.WriteString(string(word[0]))
	}
	return strings.ToUpper(str.String())
}
