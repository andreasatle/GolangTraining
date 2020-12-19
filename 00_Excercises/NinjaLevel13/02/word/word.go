package word

import (
	"strings"
)

// UseCount counts the number of occurancies of each word in the string.
func UseCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

// Count returns the number of words in the string
func Count(s string) int {
	return len(strings.Fields(s))
}
