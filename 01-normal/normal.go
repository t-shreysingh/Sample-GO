// package normal

//	func Sum(x, y int) int {
//		z := x + y
//		return z
//	}
package main

import "fmt"

func square(num int) int {
	return num * num
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	fmt.Println(square(10))
	fmt.Println(factorial(5))
}
