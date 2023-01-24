package main

import "fmt"

var pl = fmt.Println

func main() {
	statePopulations := map[string]int{
		"California": 26541651,
		"Texas":      164546456,
	}

	pl(statePopulations)

	delete(statePopulations, "Texas")
	pop, ok := statePopulations["California"]
	pl(statePopulations)
	pl(pop, ok)

	usersAge := map[string]int{
		"user1": 12,
		"user2": 13,
	}

	pl(usersAge)

}
