package main

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	name        string
	age         int
	dateOfBirth time.Time
}

func main() {
	// var aa Person
	// aa.name = "Joy"
	// aa.age = 25
	// aa.dateOfBirth = time.Now()

	// fmt.Println(aa)

	// var ab Person
	// ab.name = "Eli"
	// ab.age = 25
	// ab.dateOfBirth = time.Now()

	// fmt.Println(ab)

	// ac := Person{
	// 	name:        "cj",
	// 	age:         22,
	// 	dateOfBirth: time.Now(),
	// }

	// p2Ptr := ac
	// p2Ptr.age = 30

	// fmt.Println(ac)
	// player1 := "Asmongold"
	// player2 := "Lowko"

	// printPlayerMessage(player1)
	// printPlayerMessage(player2)

	// printPlayerNameAndAge("Joy", 25)

	// num := 1234

	// sum, err := calculateSumOfDigit(num)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(sum)

	// printPlayerInfos("Lowko")

	// num := 10
	// fmt.Println(num)

	// numCopy := num
	// fmt.Println(numCopy)

	// numCopy = 15
	// fmt.Println(num, numCopy)

	// var numPtr *int = &num

	// fmt.Println(*numPtr)

	// start, end := 10, 20
	// iterate(start, end)
	// iteratePtr(&start, &end)
	// fmt.Println(start, end)

	// mp := map[string]int{
	// 	"foo": 1,
	// }
	// updateMap(mp)
	// fmt.Println(mp)

	// vals := []int{1, 2, 3, 4, 5}

	// updateSlice(vals)
	// fmt.Println(vals)

}

func incAge(p Person) {
	p.age++
}

func updateSlice(vals []int) {
	vals = append(vals, 6)
	for i := range vals {
		vals[i] += 1
	}
	fmt.Println(vals)
}

func updateMap(mp map[string]int) {
	mp["foo"] = 43
	mp["bar"] = 56
}

func iterate(start, end int) {
	for start < end {
		fmt.Println(start)
		start++
	}
}

func iteratePtr(start, end *int) {
	for *start < *end {
		fmt.Println(*start)
		(*start)++
	}
}

func printPlayerInfos(players ...string) {
	for _, player := range players {
		fmt.Printf("Player %s is cool!\n", player)
	}
}

func calculateSumOfDigit(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("negative values are not supported")
	}

	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum, nil
}

func printPlayerMessage(player string) {
	fmt.Printf("Player %s on the starting line...", player)
	fmt.Println("The battle is about to begin!")
}

func printPlayerNameAndAge(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
