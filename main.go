package main

import (
	"github.com/yhonguevara/course-golang/webserver"
)

// "fmt"
// "runtime"

// "github.com/yhonguevara/course-golang/variables"
// "github.com/yhonguevara/course-golang/exercises"
// "github.com/yhonguevara/course-golang/iterations"
// "github.com/yhonguevara/course-golang/files"
// "github.com/yhonguevara/course-golang/functions"
// "github.com/yhonguevara/course-golang/arrays_slices"
// "github.com/yhonguevara/course-golang/maps"
// "github.com/yhonguevara/course-golang/users"
// "github.com/yhonguevara/course-golang/exercises"
// "github.com/yhonguevara/course-golang/models"
// d "github.com/yhonguevara/course-golang/defer_panic"
// "github.com/yhonguevara/course-golang/goroutines"

func main() {
	// fmt.Println("Nombre1", variables.Name)
	// variables.RestVariables()
	// fmt.Println("Nombre2", variables.Name)

	// state, textNumber := variables.ConvertToText(48715)

	// fmt.Println(state)
	// fmt.Println(textNumber)

	// if os := runtime.GOOS; os == "windows" {
	// 	fmt.Println("This is Windows")
	// } else {
	// 	fmt.Println("This is not a Windows, this is", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("This is Windows")
	// case "darwin":
	// 	fmt.Println("This is Darwin")
	// default:
	// 	fmt.Printf("This is %s \n", os)
	// }

	// number, description := exercises.Practice1("101")

	// fmt.Println(number)
	// fmt.Println(description)

	// iterations.Iteration()

	// fmt.Print(exercises.MultiplicationTable())

	// files.SaveTable()
	// files.AppendTable()
	// files.ReadFile1()
	// files.ReadFile2()

	// arrays_slices.Capacity()

	// maps.ShowMaps()

	// users.CreateUser()

	// Oscar := new(models.Man)
	// exercises.HumanBreathing(Oscar)

	// Rosmil := new(models.Women)
	// exercises.HumanBreathing(Rosmil)

	// d.ExamplePanic()

	/* channel1 := make(chan bool)
	go goroutines.MyNameSlow("Oscar Yhonnerry Guevara Siano", channel1)
	fmt.Println("Waiting for the channel")

	<-channel1*/

	webserver.Server()
}
