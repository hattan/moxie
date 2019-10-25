// Package to return allergy related information. Makes use of bitwise & . (notice allergies are 2^n)
package allergies

import "sort"

type Allergy struct {
	Name  string
	Score uint
}

var allergies = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Returns an array of Allergy sorted by Score
func getAllergies() []Allergy {
	list := make([]Allergy, len(allergies))
	for k, v := range allergies {
		list = append(list, Allergy{k, v})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Score < list[j].Score
	})
	return list
}

// Returns an array of allergy names, based on a score
func Allergies(score uint) []string {
	allergies := getAllergies()
	results := make([]string, 0)
	for _, allergy := range allergies {
		if AllergicTo(score, allergy.Name) {
			results = append(results, allergy.Name)
		}
	}
	return results
}

// Returns true if an allergy score contains the current substance
func AllergicTo(score uint, substance string) bool {
	return allergies[substance]&score > 0
}
