package day

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/templates/golang"
	"github.com/gossie/aoc-generator/util"
)

type FileData struct {
	Package    string
	Year       string
	GithubUser string
}

func CreateDay(name string) {
	err := os.Mkdir(name, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + name)
	}

	year := config.GetPropertyValue("year")
	githubUser := config.GetPropertyValue("githubUser")

	fileData := FileData{name, year, githubUser}

	dayT, err := template.New("day").Parse(golang.DayTemplate)
	if err != nil {
		panic("template could not be parsed: " + err.Error())
	}

	dayBuffer := new(bytes.Buffer)
	dayT.Execute(dayBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/%v.go", name, name), dayBuffer.String())

	dayTestT, err := template.New("dayTest").Parse(golang.DayTestTemplate)
	if err != nil {
		panic("template could not be parsed: " + err.Error())
	}

	dayTestBuffer := new(bytes.Buffer)
	dayTestT.Execute(dayTestBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/%v_test.go", name, name), dayTestBuffer.String())

	util.CreateFile(fmt.Sprintf("%v/%v.txt", name, name), "")
	util.CreateFile(fmt.Sprintf("%v/%v_test.txt", name, name), "")

	aocDayT, err := template.New("aocDayT").Parse(golang.SingleDayInAdventOfCodeTemplate)
	if err != nil {
		panic("template could not be parsed")
	}

	aocDayTBuffer := new(bytes.Buffer)
	aocDayT.Execute(aocDayTBuffer, fileData)
	content := aocDayTBuffer.String()

	mainFileContent := readMainFileContent()
	cutMainFileContent := append(make([]string, 0, len(mainFileContent)), mainFileContent[:len(mainFileContent)-1]...)
	cutMainFileContent = append(cutMainFileContent, "\n")
	cutMainFileContent = append(cutMainFileContent, content)
	cutMainFileContent = append(cutMainFileContent, "}")

	cutMainFileContent = ensureImport(cutMainFileContent, name, year, githubUser)

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

func ensureImport(cutMainFileContent []string, name, year, githubUser string) []string {
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
	return append(append(append(make([]string, 0, len(cutMainFileContent)+1), cutMainFileContent[0:endOfImport]...), fmt.Sprintf("    \"github.com/%v/adventofcode%v/%v\"", githubUser, year, name)), cutMainFileContent[endOfImport:]...)
}

func writeMainFileContent(lines []string) {
	file, err := os.OpenFile("adventofcode.go", os.O_RDWR, 0)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	for _, line := range lines {
		toWrite := line
		if toWrite != "\n" {
			toWrite += "\n"
		}
		_, err := file.WriteString(toWrite)
		if err != nil {
			panic("failed writing into file: " + err.Error())
		}
	}
	file.Sync()
}
