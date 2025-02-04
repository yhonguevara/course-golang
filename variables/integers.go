package variables

import "fmt"

func ShowIntegers() {
	var integerCommon int
	integer32 := int32(10)
	integer64 := int64(100)

	fmt.Println("Int común = ", integerCommon)
	fmt.Println("Int 32 = ", integer32)
	fmt.Println("Int 64 = ", integer64)
}