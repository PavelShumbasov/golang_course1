// Задача 1: Применение анонимной функции для суммирования

package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	result := sum(5, 4)
	fmt.Println(result)
}