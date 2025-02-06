package exercises

import (
	"strconv"
)

func Practice1(text string) (int, string) {
	number, err := strconv.Atoi(text)
	var description string

	if err != nil {
		return 0, "There was an error " + err.Error()
	}

	if number > 100 {
		description = "Is greater than 100"
	} else {
		description = "Is less than 100"
	}

	return number, description
}
