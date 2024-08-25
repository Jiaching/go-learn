package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := FizzBuzz(100)
	fmt.Println(array)
}

func FizzBuzz(num int) []string {
	var fizzBuzz []string
	for i := 1; i <= num; i++ {
		switch {
		case i%15 == 0:
			fizzBuzz = append(fizzBuzz, "FizzBuzz")
		case i%5 == 0:
			fizzBuzz = append(fizzBuzz, "Buzz")
		case i%3 == 0:
			fizzBuzz = append(fizzBuzz, "Fizz")
		default:
			fizzBuzz = append(fizzBuzz, strconv.Itoa(i))
		}
	}
	return fizzBuzz
}
