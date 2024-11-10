package main

import "fmt"

func main() {
	var factorial func(int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n - 1)
	}
	result := factorial(5)
	fmt.Println(result)
}