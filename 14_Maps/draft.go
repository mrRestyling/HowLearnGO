package main

import (
	"fmt"
)

type familyMember struct {
	Number int
	Name   string
}

func q() {

	mapByMe := make(map[int]familyMember)

	fmt.Println(mapByMe)

	mapByMe[1] = familyMember{
		78126661488,
		"Egor",
	}

	fmt.Println(mapByMe)

	// delete(mapByMe, 1)
	// fmt.Println(mapByMe)

	family := []familyMember{
		{
			915,
			"Sweta",
		},
		{
			916,
			"Irina",
		},
		{
			213,
			"Egor",
		},
	}

	for _, people := range family {
		if _, ok := mapByMe[people.Number]; !ok {
			mapByMe[people.Number] = people
		}
	}

	fmt.Println(mapByMe)
}
