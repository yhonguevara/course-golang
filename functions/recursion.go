package functions

import "fmt"

func Exponence(value int) {
	if value > 10000000 {
		return
	}

	fmt.Println(value)
	Exponence(value * 2)
}
