package java

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/templates/java"
	"github.com/gossie/aoc-generator/util"
)

type fileData struct {
	Day        int
	Year       string
	GithubUser string
}

func CreateDay(day int) {
	dayPackage := fmt.Sprintf("day%d", day)
	year := config.GetPropertyValue("year")
	githubUser := config.GetPropertyValue("githubUser")

	util.CreateDirectory(fmt.Sprintf("src/main/java/adventofcode/%v", dayPackage))
	util.CreateDirectory(fmt.Sprintf("src/test/java/adventofcode/%v", dayPackage))
	util.CreateDirectory(fmt.Sprintf("src/main/resources/adventofcode/%v", dayPackage))
	util.CreateDirectory(fmt.Sprintf("src/test/resources/adventofcode/%v", dayPackage))

	fileData := fileData{day, year, githubUser}

	dayT, err := template.New("day").Parse(java.DayTemplate)
	if err != nil {
		panic("template could not be parsed: " + err.Error())
	}

	dayBuffer := new(bytes.Buffer)
	dayT.Execute(dayBuffer, fileData)
	util.CreateFile(fmt.Sprintf("src/main/java/adventofcode/%v/Day%d.java", dayPackage, day), dayBuffer.String())

	dayTestT, err := template.New("dayTest").Parse(java.DayTestTemplate)
	if err != nil {
		panic("template could not be parsed: " + err.Error())
	}

	dayTestBuffer := new(bytes.Buffer)
	dayTestT.Execute(dayTestBuffer, fileData)
	util.CreateFile(fmt.Sprintf("src/test/java/adventofcode/%v/Day%dTest.java", dayPackage, day), dayTestBuffer.String())

	util.CreateFile(fmt.Sprintf("src/main/resources/adventofcode/%v/%v.txt", dayPackage, dayPackage), "")
	util.CreateFile(fmt.Sprintf("src/test/resources/adventofcode/%v/%v.txt", dayPackage, dayPackage), "")

	aocDayT, err := template.New("aocDayT").Parse(java.SingleDayInAdventOfCodeTemplate)
	if err != nil {
		panic("template could not be parsed")
	}

	aocDayTBuffer := new(bytes.Buffer)
	aocDayT.Execute(aocDayTBuffer, fileData)
	content := aocDayTBuffer.String()

	mainFileContent := readMainFileContent()
	cutMainFileContent := append(make([]string, 0, len(mainFileContent)), mainFileContent[:len(mainFileContent)-3]...)
	cutMainFileContent = append(cutMainFileContent, "\n")
	cutMainFileContent = append(cutMainFileContent, content)
	cutMainFileContent = append(cutMainFileContent, "    }")
	cutMainFileContent = append(cutMainFileContent, "}")

	writeMainFileContent(cutMainFileContent)

	fmt.Println("Created files", dayPackage)
}

func readMainFileContent() []string {
	file, err := os.Open("src/main/java/adventofcode/AdventOfCode.java")
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

func writeMainFileContent(lines []string) {
	file, err := os.OpenFile("src/main/java/adventofcode/AdventOfCode.java", os.O_RDWR, 0)
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
