package main

import "fmt"

func main() {
	m := map[string][]string{}
	m[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	m[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range m {
		fmt.Println(k)
		for i, w := range v {
			fmt.Printf("\t%d\t%s\n", i, w)
		}
	}

}
