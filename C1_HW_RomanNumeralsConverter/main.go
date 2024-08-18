package main

import (
	"fmt"
)

func main() {
	//fmt.Println(IsCorrect(romanToInt("CIV"), 104))
	//fmt.Println(IsCorrect(romanToInt("LVIII"), 58))
	fmt.Println(IsCorrect(romanToInt("MCMXCIV"), 1994))
	//fmt.Println(IsCorrect(romanToInt("MCM"), 1900))
}

func IsCorrect(value int, expect int) bool {
	return value == expect
}

func romanToInt(romanInput string) int {
	romanNumberMap := make(map[string]int)
	romanNumberMap["M"] = 1000
	romanNumberMap["D"] = 500
	romanNumberMap["C"] = 100
	romanNumberMap["L"] = 50
	romanNumberMap["X"] = 10
	romanNumberMap["V"] = 5
	romanNumberMap["I"] = 1

	romanNumberReplaceMap := make(map[string]int)
	romanNumberReplaceMap["CM"] = 900
	romanNumberReplaceMap["CD"] = 400
	romanNumberReplaceMap["XC"] = 90
	romanNumberReplaceMap["XL"] = 40
	romanNumberReplaceMap["IX"] = 9
	romanNumberReplaceMap["IV"] = 4

	// 建立一個Stack, 先進先出 ，預期會是 從小到大
	// 從右到左去輸入丟到Stack
	// 將Top 與接下來的數字比對，發現可以被取代，就將數值取代
	// 最後將stack 加總

	romanInputIndex := 0
	lastRomanInput := string(romanInput[len(romanInput)-1])
	charNumber, exists := romanNumberMap[lastRomanInput]
	if !exists {
		fmt.Println("請重新輸入，字元不符合")
		return 0
	}

	nums := make([]int, 0)
	nums = append(nums, charNumber)
	fmt.Println(nums)
	romanInputIndex++
	for index := romanInputIndex; index < len(romanInput); index++ {
		currentRomanInput := string(romanInput[len(romanInput)-index-1])
		fmt.Println("Current Roman Input ", currentRomanInput)
		replaceRoman := currentRomanInput + lastRomanInput
		fmt.Println(replaceRoman)
		replaceValue, exists := romanNumberReplaceMap[replaceRoman]
		if exists {
			nums[len(nums)-1] = replaceValue
			fmt.Println(nums)
			lastRomanInput = currentRomanInput
			continue
		}

		currentRomanNumber, exists := romanNumberMap[currentRomanInput]
		if !exists {
			fmt.Println("請重新輸入，字元不符合")
			return 0
		}

		nums = append(nums, currentRomanNumber)
		lastRomanInput = currentRomanInput
		fmt.Println(nums)
	}

	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println(sum)
	return sum
}
