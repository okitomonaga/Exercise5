package main

import "fmt"
import "strconv"

func main() {
	var a string
	var b string
	var c string

	fmt.Scan(&a, &b, &c)
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	fmt.Println(ai)
	fmt.Println(bi)
	fmt.Println(c)
}
