package exercises

import (
	"fmt"

	"github.com/yhonguevara/course-golang/interfaces"
)

func HumanBreathing(hb interfaces.Human) {
	hb.Breathe()

	fmt.Printf("Soy un/a %s, y estoy respirando \n", hb.Sex())
}