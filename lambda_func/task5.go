// Фильтрация слайса

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var evens []int
	for _, num := range numbers {
		if func(n int) bool { return n % 2 == 0 }(num) {
			evens = append(evens, num)
		}
	}

	fmt.Println(evens)
}