package functions

import "fmt"

func Calculations() {
	addition := func(number1 int, number2 int) int {
		return number1 + number2
	}

	fmt.Println(addition(10, 32))
}
