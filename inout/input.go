package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var num int = 43
	// fmt.Println("hello world")
	// fmt.Println("hello", "world", "!", num)
	// fmt.Print("hello", "world", "!", num)
	// 	fmt.Println(
	// 		`Hello, my name is joy.
	// I like to code.`)

	// var name string = "joy"
	// var age int = 25
	// var height int = 164
	// fmt.Printf("Hello and welcome %s\nYour age is %d and your height is %d\nGood for you!\n", name, age, height)

	// reader := bufio.NewReader(os.Stdin)
	// line, err := reader.ReadString('\n')

	// if err != nil {
	// 	panic(err)
	// }

	// line = strings.TrimSpace(line)
	// fmt.Println("Line 1:", line)

	// line, err = reader.ReadString('\n')

	// if err != nil {
	// 	panic(err)
	// }

	// line = strings.TrimSpace(line)
	// fmt.Println("Line 2:", line)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(strings.TrimSpace(line))

	if err != nil {
		panic(err)
	}

	fmt.Println("My number is:", num)
	fmt.Println("My number is:", num+23)

}
