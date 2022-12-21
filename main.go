package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "create":
		if len(os.Args) == 3 {
			createDirectory(os.Args[2])
		} else {
			println("I will help you later, I promise")
		}
	case "init":
		if len(os.Args) == 3 {
			initializeYear(os.Args[2])
		} else {
			println("I will help you later, I promise")
		}
	case "--help":
		printHelp()
	default:
		printHelp()
	}
}

func initializeYear(year string) {
	directoryName := "advent-of-code-" + year
	err := os.Mkdir(directoryName, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + directoryName)
	}

	createFile(fmt.Sprintf("%v/go.mod", directoryName), fmt.Sprintf("module github.com/gossie/adventofcode%v\n\ngo 1.19\n", year))
	createFile(fmt.Sprintf("%v/adventofcode.go", directoryName), "package main\n\nimport (\n    \"fmt\"\n    \"time\"\n)\n\nfunc main() {\n\n}\n")

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

	mainFileContent := readMainFileContent()
	cutMainFileContent := append(make([]string, 0, len(mainFileContent)), mainFileContent[:len(mainFileContent)-2]...)
	cutMainFileContent = append(cutMainFileContent, "\n")
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    fmt.Println(\"\\nPerforming tasks of %v\")", name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    start%vPart1 := time.Now().UnixMilli()", name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    %vPart1 := %v.Part1(\"%v/%v_test.txt\")", name, name, name, name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    fmt.Println(\"%v, task 1: \", %vPart1, \", took\", (time.Now().UnixMilli() - start%vPart1), \"ms\")", name, name, name))
	cutMainFileContent = append(cutMainFileContent, "\n")
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    fmt.Println(\"\\nPerforming tasks of %v\")", name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    start%vPart2 := time.Now().UnixMilli()", name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    %vPart2 := %v.Part2(\"%v/%v_test.txt\")", name, name, name, name))
	cutMainFileContent = append(cutMainFileContent, fmt.Sprintf("    fmt.Println(\"%v, task 2: \", %vPart2, \", took\", (time.Now().UnixMilli() - start%vPart2), \"ms\")", name, name, name))
	cutMainFileContent = append(cutMainFileContent, "}")
	cutMainFileContent = append(cutMainFileContent, "\n")

	cutMainFileContent = ensureImport(cutMainFileContent, name)

	writeMainFileContent(cutMainFileContent)

	fmt.Println("Created files", name)
}

func readMainFileContent() []string {
	file, err := os.Open("adventofcode.go")
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ensureImport(cutMainFileContent []string, name string) []string {
	foundImportStart := false
	var endOfImport int
	for i, line := range cutMainFileContent {
		if line == "import (" {
			foundImportStart = true
		}
		if foundImportStart {
			if line == ")" {
				endOfImport = i
			}
		}
	}
	return append(append(append(make([]string, 0, len(cutMainFileContent)+1), cutMainFileContent[0:endOfImport]...), fmt.Sprintf("    \"github.com/gossie/adventofcode2023/%v\"", name)), cutMainFileContent[endOfImport:]...)
}

func writeMainFileContent(lines []string) {
	file, err := os.OpenFile("adventofcode.go", os.O_RDWR, 0)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic("failed writing into file: " + err.Error())
		}
	}
	file.Sync()
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

func printHelp() {
	println("I will help you later, I promise")
}
