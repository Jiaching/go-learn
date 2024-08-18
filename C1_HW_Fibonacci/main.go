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
	// 先設定slice 大小，採用類似array存取與設定數值
	//numOfFibonacci := make([]int, num)
	//numOfFibonacci[0] = 1
	//numOfFibonacci[1] = 1
	//for index := 2; index < num; index++ {
	//	numOfFibonacci[index] = numOfFibonacci[index-2] + numOfFibonacci[index-1]
	//}

	// 先設定初始slice 長度，後續透過 append 新增數字到slice 後面
	var numOfFibonacci = []int{1, 1}
	for index := 2; index < num; index++ {
		numOfFibonacci = append(numOfFibonacci, numOfFibonacci[index-2]+numOfFibonacci[index-1])
	}

	return numOfFibonacci
}
// 第二種寫法，將slice 傳遞參考到func
func secondMainToFibonacci() {
	var num int
	var numOfFibonacci = []int{1, 1}

	fmt.Println("請輸入要計算費式數列的數字")
	fmt.Scanln(&num)

	secondCalculateFibonacci(num, &numOfFibonacci)
	if numOfFibonacci == nil {
		fmt.Println("請輸入大於1的數字")
		return
	}

	fmt.Println(numOfFibonacci)
}

func secondCalculateFibonacci(num int, numOfFibonacci *[]int) {
	if num < 2 {
		numOfFibonacci = nil
		return
	}
	for index := len(*numOfFibonacci); index < num; index++ {
		*numOfFibonacci = append(*numOfFibonacci, (*numOfFibonacci)[index-2]+(*numOfFibonacci)[index-1])
	}
}
