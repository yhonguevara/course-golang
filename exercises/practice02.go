package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() string {
	scannner := bufio.NewScanner(os.Stdin)

	var number int
	var err error
	var output string

	for {
		fmt.Print("Ingresa un número para construir la tabla de multiplicar: ")

		if scannner.Scan() {
			number, err = strconv.Atoi(scannner.Text())
			if err != nil {
				// fmt.Println("The date entry is incorrect " + err.Error())
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		output += fmt.Sprintf("%d x %d = %d \n", number, i, (number * i))
	}

	return output
}
