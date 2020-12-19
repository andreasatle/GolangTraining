package dog

import (
	"fmt"
	"testing"
)

func TestYear(t *testing.T) {
	manYears := 10
	dogYears := Years(manYears)
	if dogYears != 7*manYears {
		t.Errorf("Got: %v, Expected: %v\n", dogYears, 7*manYears)
	}
}

func TestYearTwo(t *testing.T) {
	manYears := 10
	dogYears := YearsTwo(manYears)
	if dogYears != 7*manYears {
		t.Errorf("Got: %v, Expected: %v\n", dogYears, 7*manYears)
	}
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := 0; y < 100; y++ {
			Years(y)
		}
	}
}

func BenchmarkYearTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := 0; y < 100; y++ {
			YearsTwo(y)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// output:
	// 70
}
