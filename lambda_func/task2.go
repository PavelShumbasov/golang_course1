// Напиши программу, которая возвращает функцию из анонимной функции. 
// Функция должна умножать число на переданный множитель.

package main

import "fmt"

func main() {
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}
	double := multiplier(2) // сделали функцию, которая будет умножать на 2
	result := double(5) 	// умножаем на 2 пятерку
	fmt.Println(result)
}