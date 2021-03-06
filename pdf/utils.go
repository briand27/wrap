package pdf

import (
	"unicode"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getHeight(lines []styledLine) int {
	height := 0
	for _, line := range lines {
		height += line.leading() + 1
	}
	return height
}

func removeSpaceBefore(s string) string {
	// Can't remove from something empty...
	if s == "" {
		return ""
	}

	i := 0
	for currentlyspace := true; i < len(s) && currentlyspace; i++ {
		currentlyspace = unicode.IsSpace(rune(s[i]))
	}

	// Now i is the point where the first nonspace is, thus
	return s[i-1:]
}

func removeSpaceAfter(s string) string {
	// Can't remove from something empty...
	if s == "" {
		return ""
	}

	s = reverse(s)
	s = removeSpaceBefore(s)
	s = reverse(s)
	return s
}

func reverse(s string) string {
	c := []rune(s)

	for i, j := 0, len(c)-1; i < len(c)/2; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	return string(c)
}
