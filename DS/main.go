package main

import (
	"fmt"
)

func main() {
	// zombies := []string{"Paul", "Gooey", "George", "Lucy"}
	// fmt.Println(zombies)

	// zombies[1] = "Katya"

	// for i := 0; i < len(zombies); i++ {
	// 	fmt.Println(zombies[i])
	// }

	// number := []int{1, 2, 3, 4}

	// for i := 0; i < len(number); i++ {
	// 	number[i]++
	// }

	// fmt.Println(number)

	// numSlice := []int{1, 2, 3, 4}

	// for i, _ := range numSlice {
	// 	// fmt.Println("Incrementing", element)
	// 	numSlice[i] += 1
	// }

	// fmt.Println(numSlice)

	// zombies := []string{}
	// zombies = append(zombies, "Paul", "George", "Katya")
	// zombies1 := append(zombies, "Lucy")
	// fmt.Println(zombies)
	// fmt.Println(zombies1)

	// odds := []int{1, 3, 5, 7}
	// evens := []int{2, 4, 6, 8}

	// result := append(odds, evens...)
	// fmt.Println(result)

	// numbers := []int{1, 2, 3, 4, 5, 6}

	// subslice := numbers[2:]

	// fmt.Println(subslice)

	// zombies := []string{"Paul", "Katya", "George", "Lucy"}
	// toRemove := 2

	// zombies = append(zombies[:toRemove], zombies[toRemove+1:]...)

	// fmt.Println(zombies)

	// zombies := [2]string{"Paul"}
	// fmt.Println(zombies)

	playerHPs := map[string]int{
		"coolboy": 10,
		"hotboy":  50,
		"tallboy": 30,
	}

	fmt.Println(playerHPs)
	fmt.Println(playerHPs["tallboy"])
}
