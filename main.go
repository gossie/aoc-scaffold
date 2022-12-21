package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "create":
		createDirectory(os.Args[2])
	default:
		panic("unknown command: " + command)
	}
}

func createDirectory(name string) {
	err := os.Mkdir(name, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + name)
	}

	createFile(fmt.Sprintf("%v/%v.go", name, name))
	createFile(fmt.Sprintf("%v/%v_test.go", name, name))
	createFile(fmt.Sprintf("%v/%v.txt", name, name))
	createFile(fmt.Sprintf("%v/%v_test.txt", name, name))
}

func createFile(name string) {
	testInput, err := os.Create(name)
	if err != nil {
		panic("failed to create file [" + name + "]: " + err.Error())
	}
	defer testInput.Close()
}
