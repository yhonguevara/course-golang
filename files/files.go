package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yhonguevara/course-golang/exercises"
)

var fileName string = "./files/txt/table.txt"

func SaveTable() {
	output := exercises.MultiplicationTable()
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file " + err.Error())
		return
	}

	fmt.Fprintln(file, output)
	file.Close()
}

func AppendTable() {
	output := exercises.MultiplicationTable()

	if !Append(fileName, output) {
		fmt.Println("Error concatenating content")
	}
}

func Append(fileName string, output string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error during Append " + err.Error())
		return false
	}

	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error during WriteString " + err.Error())
		return false
	}

	file.Close()

	return true
}

func ReadFile1() {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file " + err.Error())
	}

	fmt.Println(string(file))
}

func ReadFile2() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("There was an error reading the file " + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record := scanner.Text()
		fmt.Println("> " + record)
	}

	file.Close()
}
