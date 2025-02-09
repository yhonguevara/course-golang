package arrays_slices

import "fmt"

func ShowArrays() {
	var table [10]int = [10]int{10, 2, 59}
	var matrix [10][20]int

	table[7] = 33
	table[2] = 54

	fmt.Println(table)

	table2 := [10]string{"Hola", "Mundo", "!"}

	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}

	matrix[5][10] = 20
	fmt.Println(matrix)
}
