// package user contains sorted user information
package user

import (
	"fmt"
	"strings"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// String implements the Stringer interface
func (c User) String() string {
	str := fmt.Sprintf("%v %v, %v years old\n\t", c.First, c.Last, c.Age)

	return str + strings.Join(c.Sayings, "\n\t") + "\n"
}

// ByAge implements the sort.Interface
type ByAge []User

// Len is part of sort.Interface.
func (s ByAge) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface.
func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// ByLast implements the sort.Interface
type ByLast []User

// Len is part of sort.Interface.
func (s ByLast) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s ByLast) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface.
func (s ByLast) Less(i, j int) bool {
	return s[i].Last < s[j].Last
}
