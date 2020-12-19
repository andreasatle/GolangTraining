// Package dog converts human-years into dog-years
//
// copyright: Andreas Atle, atle.andreas@gmail.com
package dog

// Years compute the dog-years from human-years.
func Years(humanYears float64) float64 {
	dogYears := 7.0 * humanYears
	return dogYears
}
