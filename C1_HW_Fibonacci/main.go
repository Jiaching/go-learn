package main

import "fmt"

func main() {
	var num int

	fmt.Println("請輸入要計算費式數列的數字")
	fmt.Scanln(&num)

	fibonacci := calculateFibonacci(num)
	if fibonacci == nil {
		fmt.Println("請輸入大於1的數字")
		return
	}

	fmt.Println(fibonacci)
}

func calculateFibonacci(num int) []int {
	if num < 2 {
		return nil
	}

	var numOfFibonacci = []int{1, 1}
	for index := len(numOfFibonacci); index < num; index++ {
		numOfFibonacci = append(numOfFibonacci, numOfFibonacci[index-2]+numOfFibonacci[index-1])
	}

	return numOfFibonacci
}
