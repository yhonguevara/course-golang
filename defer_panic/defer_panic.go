package defer_panic

import (
	"fmt"
	"log"
)

func ViewDefer() {
	fmt.Println("Primer Mendaje")
	defer fmt.Println("Último Mensaje")
	fmt.Println("Segundo Mensaje")
}

func ExamplePanic() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered in ExamplePanic \n %v", r)
		}
	}()
	panic("Example Panic")
}
