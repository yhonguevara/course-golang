package arrays_slices

import "fmt"

func ShowSlice() {
	var table []int = []int{2, 5, 9}
	var array [10]int = [10]int{54, 564, 321, 398, 69, 362, 12, 7}

	fmt.Println(table)

	slice1 := array[3:]  // Slice creado desde un vector, desde la posición 3
	slice2 := array[:5]  // Slice creado desde un vector, desde la posición 0 hasta la 5
	slice3 := array[6:7] // Slice creado desde un vector desde la posición 6 hasta la 7

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Length %d, Capacity %d \n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i <= 99; i++ {
		nums = append(nums, i+1)
	}

	fmt.Println(nums)
	fmt.Printf("Length %d, Capacity %d \n", len(nums), cap(nums))
}
