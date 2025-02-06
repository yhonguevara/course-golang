package iterations

import "fmt"

func Iteration() {
	// for {
	// 	fmt.Println("Hola!")
	// 	break
	// }

	// i := 0

	// for i < 10 {
	// 	fmt.Println(i + 1)
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i + 1)
	// }

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}

}
