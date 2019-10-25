// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

// Hey should have a comment documenting it.
func Hey(remark string) string {
	last := remark[len(remark)-1:]
	switch last {
	case ".":
		return "Whatever."
	case "!":
		return "Whoa, chill out!"
	case "?":
		return "Sure."
	default:
		return "Whoa, chill out!"
	}

	return ""
}
