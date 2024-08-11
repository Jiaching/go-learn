package main

import "fmt"

func main() {
	var characterJob string
	var characterPersonalAdjective string
	var characterGoal string
	var characterSituationAdverb string
	var characterVerb string

	fmt.Println("歡迎來到故事模板，接下來會請你輸入不同的詞彙，來透過模板創建一個故事")
	// 隨機輸入名詞/動詞/形容詞
	Prompt("請輸入角色現況的職業(e.g.初學者/冒險家/流浪漢)", &characterJob)
	Prompt("請輸入角色的人格特質(e.g.勇者/膽小/堅強/衝動/隨便)", &characterPersonalAdjective)
	Prompt("請輸角色未來想要成為的目標(e.g. 勇者/獵人/富翁)", &characterGoal)
	Prompt("請輸入角色今日的狀態(e.g.幸運地/糟糕地/幸輕鬆地/艱難地)", &characterSituationAdverb)
	Prompt("請輸入角色今日的動作(e.g. 發現/擊殺/收服/催眠)", &characterVerb)

	fmt.Println("\n這是屬於您的故事")
	story := fmt.Sprintf("從前有一位%s，個性非常的%s\n"+
		"由於想要成為一位%s，因此開始踏上冒險的旅途\n"+
		"%s，今日%s怪物史萊姆，成為初心者未來的夥伴，繼續冒險", characterJob, characterPersonalAdjective, characterGoal, characterSituationAdverb, characterVerb)
	fmt.Println(story)

	// 從前有一位初心者(名詞)，個性非常的勇敢(形容詞)
	// 由於想要成為一位勇者(名詞)，因此開始踏上冒險的旅途
	// 幸運地(副詞)，今日遇到(動詞)怪物史萊姆，成為初心者未來的夥伴，繼續冒險

}

func Prompt(prompt string, input *string) {
	fmt.Println(prompt)
	fmt.Scan(input)
}
