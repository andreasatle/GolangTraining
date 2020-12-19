// Custom sort
package main

import (
	"fmt"
	"sort"

	"github.com/andreasatle/Udemy/GolangTraining/00_Excercises/NinjaLevel08/05/user"
)

func main() {
	u1 := user.User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user.User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user.User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user.User{u1, u2, u3}

	fmt.Println(users)
	// your code goes here
	sort.Sort(user.ByAge(users))
	fmt.Println(users)

	sort.Sort(user.ByLast(users))
	fmt.Println(users)

}
