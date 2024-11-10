/*
Вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(n uint) uint {
		numStr := strconv.Itoa(int(n))
		var resultStr string

		for _, ch := range numStr {
			if ch == '0' || (ch - '0') % 2 != 0 {
				continue
			}
			resultStr += string(ch)
		}
		if resultStr == "" {return 100}
		result, _ := strconv.Atoi(resultStr)
		return uint(result)
	}
	fmt.Println(fn(727178))
}