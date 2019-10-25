// Let's detect some similar words
package anagram

import (
	"sort"
	"strings"
)

// Sort a string by converting it to an array of characters first
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Detect converts both the subject and each candidate into lower case and sorts by characters. Then compares for equality
func Detect(subject string, candidates []string) []string {
	results := make([]string, 0)
	subject = strings.ToLower(subject)
	for _, candidate := range candidates {
		loweredCandidate := strings.ToLower(candidate)
		if loweredCandidate == subject {
			return results
		}
		sorted := SortString(subject)
		sortedCandidate := SortString(loweredCandidate)
		if sorted == sortedCandidate {
			results = append(results, candidate)
		}
	}
	return results
}
