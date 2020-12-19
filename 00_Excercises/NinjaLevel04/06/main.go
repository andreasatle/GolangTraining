package main

import "fmt"

func main() {
	states := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}

	print(states)

	// Append and remove an element to see what happens with cap...
	states = append(states, "Sweden")
	states = states[:50]

	print(states)
}

func print(states []string) {
	for i := 0; i < len(states); i++ {
		if i%5 == 0 {
			fmt.Println(i, states[i])
		}
	}
	fmt.Println("Len:", len(states))
	fmt.Println("Cap:", cap(states))
}
