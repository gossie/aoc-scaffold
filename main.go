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
	case "init":
		initializeYear(os.Args[2])
	case "--help":
		println("I will help you later, I promise")
	default:
		panic("unknown command: " + command)
	}
}

func initializeYear(year string) {
	directoryName := "advent-of-code-" + year
	err := os.Mkdir(directoryName, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + directoryName)
	}

	createFile(fmt.Sprintf("%v/go.mod", directoryName), fmt.Sprintf("module github.com/gossie/adventofcode%v\n\ngo 1.19\n", year))
	createFile(fmt.Sprintf("%v/adventofcode.go", directoryName), "package main\n\nfunc main() {\n\n}\n")

	fmt.Println("Created new project for the advent of code", year)
}

func createDirectory(name string) {
	err := os.Mkdir(name, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + name)
	}

	createFile(fmt.Sprintf("%v/%v.go", name, name), fmt.Sprintf("package %v\n\nfunc Part1(filename string) int {\n    return 0;\n}\n\nfunc Part2(filename string) int {\n    return 0;\n}\n", name))
	createFile(fmt.Sprintf("%v/%v_test.go", name, name), fmt.Sprintf("package %v_test\n\nimport (\n    \"testing\"\n\n    \"github.com/gossie/adventofcode2023/%v\"\n)\n\nfunc TestPart1(t *testing.T) {\n    part1 := %v.Part1(\"%v_test.txt\")\n    if part1 != 0 {\n        t.Fatalf(\"part1 = %%v\", part1)\n    }\n}\n\nfunc TestPart2(t *testing.T) {\n    part2 := %v.Part2(\"%v_test.txt\")\n    if part2 != 0 {\n        t.Fatalf(\"part2 = %%v\", part2)\n    }\n}\n", name, name, name, name, name, name))
	createFile(fmt.Sprintf("%v/%v.txt", name, name), "")
	createFile(fmt.Sprintf("%v/%v_test.txt", name, name), "")

	fmt.Println("Created files", name)
}

func createFile(name, content string) {
	testInput, err := os.Create(name)
	if err != nil {
		panic("failed to create file [" + name + "]: " + err.Error())
	}

	_, err = testInput.WriteString(content)
	if err != nil {
		panic("failed to wirte into file [" + name + "]: " + err.Error())
	}
	testInput.Sync()

	defer testInput.Close()
}
