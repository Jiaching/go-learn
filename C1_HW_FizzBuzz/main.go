package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := FizzBuzzBetterMemory(100)
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

func FizzBuzzBetterMemory(num int) []string {
	var fizzBuzz = make([]string, num)
	for i := 1; i <= num; i++ {
		switch {
		case i%15 == 0:
			fizzBuzz[i-1] = "FizzBuzz"
		case i%5 == 0:
			fizzBuzz[i-1] = "Buzz"
		case i%3 == 0:
			fizzBuzz[i-1] = "Fizz"
		default:
			fizzBuzz[i-1] = strconv.Itoa(i)
		}
	}

	return fizzBuzz
}
