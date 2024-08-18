### 目標:

在第一個挑戰中，您將撰寫程式來從某個數字開始計算 Fibonacci 數列。 您會撰寫一個函式，傳回包含斐波那契序列中所有數字的切片。 序列會根據使用者輸入的數位來執行計算的結果。 輸入編號必須大於兩個。 讓我們假設小於 2 的數字將產生錯誤，並傳回 nil 配量。

請記住，Fibonacci 數列是一個數字清單，其中每個數字都是前兩個 Fibonacci 數字的總和。 例如，6 的數字序列是 1,1,2,3,5,8、7 的數字序列是 1,1,2,3,5,8,13、8 的數字序列是 1,1,2,3,5,8,13,21，依此類推。
### 要求:

### 執行結果輸出
**正常**
```
請輸入要計算費式數列的數字
20
[1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765]

```
**例外處理**

輸入的數字需要大於1
```
請輸入要計算費式數列的數字
1
請輸入大於1的數字
Exiting.

```

### Notes
- 原本想要透過`*numOfFibonacci[0]` 結果出現錯誤`'*[]int' does not support indexing`，因此Google後調整寫法
```go
func secondCalculateFibonacci(num int, numOfFibonacci *[]int) {
//error: type '*[]int' does not support indexing)
*numOfFibonacci[0] = 0
```
- Go Lang slice append 操作的記憶體方式
  1. Append 後獲得的slice 與原傳入slice 指向相同array的指標
  2. 因此原slice 與新的slice 各自append 可能會互相改到數值
  3. 要透過新的append 建立新的自己，才能指向兩個不同array的指標
  
