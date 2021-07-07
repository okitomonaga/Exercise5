package main

import "fmt"
import "strconv"

func calcPlus(x int, y int) int {
	return x + y
}

func calcMinus(x int, y int) int {
	return x - y
}

func main() {
	var a string
	var b string
	var c string
	var result int

	fmt.Scan(&a, &b, &c)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	switch c {
	case "+":
		fmt.Println("+が入力されました")
		result = calcPlus(ai, bi)
	case "-":
		fmt.Println("-が入力されました") // 2つめが出力される
		result = calcMinus(ai, bi)
	case "*":
		fmt.Println("*が入力されました") // 2つめが出力される
	case "/":
		fmt.Println("/が入力されました") // 2つめが出力される
	default:
		fmt.Println("エラー")
	}

	// 	fmt.Println(ai)
	// 	fmt.Println(bi)
	// 	fmt.Println(c)
	fmt.Println(result)
}
