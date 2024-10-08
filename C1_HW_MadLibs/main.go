package main

import (
	"fmt"
	"strconv"
	"strings"
)

const EmptyString = ""

// 每一個詞彙都有3個簡單的提示，玩家只需要輸入數字，選擇提示的詞彙，或是自行輸入想要的詞彙
// 需要判斷玩家是否輸入為數字或是非數字/非小數點的字串
// 提示範例： 請輸入XX(名詞) 或是輸入數字，使用提示的詞彙(1. 2. 3. 4...)

func main() {
	var characterJob string
	var characterPersonalAdjective string
	var characterGoal string
	var characterSituationAdverb string
	var characterVerb string

	var jobHints = []string{"初學者", "冒險家", "流浪漢"}
	var personalAdjectiveHints = []string{"勇者", "膽小", "堅強", "衝動", "隨便"}
	var goalHints = []string{"勇者", "獵人", "富翁"}
	var situationHints = []string{"幸運地", "糟糕地", "幸輕鬆地", "艱難地"}
	var verbHints = []string{"發現", "擊殺", "收服", "催眠"}
	fmt.Println("歡迎來到故事模板，接下來會請你輸入不同的詞彙，來透過模板創建一個故事")
	// 隨機輸入名詞/動詞/形容詞
	PromptHints("請輸入角色現況的職業 %s", jobHints, &characterJob)
	PromptHints("請輸入角色的人格特質 %s", personalAdjectiveHints, &characterPersonalAdjective)
	PromptHints("請輸角色未來想要成為的目標 %s", goalHints, &characterGoal)
	PromptHints("請輸入角色今日的狀態 %s", situationHints, &characterSituationAdverb)
	PromptHints("請輸入角色今日的動作 %s", verbHints, &characterVerb)

	fmt.Println("\n這是屬於您的故事")
	story := fmt.Sprintf("從前有一位%s，個性非常的%s\n"+
		"由於想要成為一位%s，因此開始踏上冒險的旅途\n"+
		"%s，今日%s怪物史萊姆，成為%s未來的夥伴，繼續冒險", characterJob, characterPersonalAdjective, characterGoal, characterSituationAdverb, characterVerb, characterJob)
	fmt.Println(story)
}

func Prompt(prompt string, input *string) {
	fmt.Println(prompt)
	fmt.Scanln(input)
}

// 提示範例： 請輸入XX(名詞) 或是輸入數字，使用提示的詞彙(1.aaa 2.bbbb 3.ccc)
// 回傳玩家輸入的詞彙
func PromptHints(prompt string, hints []string, userInput *string) {
	var hintsBuilder strings.Builder

	// 組合提示詞
	for i, hint := range hints {
		var number = i + 1
		hintsBuilder.WriteString(fmt.Sprintf("%d. %s ", number, hint))
	}

	var promptWithHints = fmt.Sprintf(prompt, hintsBuilder.String())
	for {
		Prompt(promptWithHints, userInput)
		number, err := strconv.Atoi(*userInput)

		// 表示玩家自行輸入字串
		if err != nil {
			*userInput = strings.TrimSpace(*userInput)
			if *userInput == EmptyString {
				fmt.Println("輸入字串為空，請重新輸入")
				continue
			}
			break
		}

		var hintIndex = number - 1
		if hintIndex < 0 || hintIndex >= len(hints) {
			fmt.Println("輸入數字錯誤，請重新輸入")
			continue
		}

		// 取的陣列內提示的詞
		if hintIndex < len(hints) {
			*userInput = hints[hintIndex]
			break
		}

		break
	}
}
