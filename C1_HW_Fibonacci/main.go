package main

import "fmt"

func main() {
	var num int
	var numOfFibonacci = []int{1, 1}
	fmt.Println("請輸入要計算費式數列的數字")
	fmt.Scanln(&num)
	if num < 2 {
		fmt.Println("請輸入大於1的數字")
		return
	}
	for index := len(numOfFibonacci); index < num; index++ {
		numOfFibonacci = append(numOfFibonacci, numOfFibonacci[index-2]+numOfFibonacci[index-1])
	}

	fmt.Println(numOfFibonacci)
}

//func fibonacci(num int, numOfFibonacci *[]int) {
//	for index := len(*numOfFibonacci); index < num; index++ {
//		*numOfFibonacci= append(*numOfFibonacci, numOfFibonacci[index-2], numOfFibonacci[index-1])
//	}
//}
