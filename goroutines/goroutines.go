package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MyNameSlow(name string, channel1 chan bool) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1 * time.Second)
		fmt.Print(letter)
	}

	fmt.Println()
	channel1 <- true
}
