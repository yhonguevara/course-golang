package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func EnterNumber() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter number 1:")

	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("The data entered is incorrect " + err.Error())
		}
	}
}
