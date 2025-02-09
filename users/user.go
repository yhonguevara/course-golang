package users

import (
	"fmt"
	"time"

	"github.com/yhonguevara/course-golang/models"
)

func CreateUser() {
	user := new(models.User)

	user.AddUser(1, "Oscar", time.Now(), true)

	fmt.Println(user)
}
