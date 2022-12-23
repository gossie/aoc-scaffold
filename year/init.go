package year

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/templates/golang"
	"github.com/gossie/aoc-generator/util"
)

type fileData struct {
	GithubUser string
	Year       string
}

func InitializeYear(year string, githubUser string) {
	directoryName := "advent-of-code-" + year
	err := os.Mkdir(directoryName, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + directoryName)
	}

	config.WriteConfig(map[string]string{"year": year, "githubUser": githubUser})

	fileData := fileData{githubUser, year}

	goModT, err := template.New("goMod").Parse(golang.GoModTemplate)
	if err != nil {
		panic("template could not be parsed")
	}

	goModBuffer := new(bytes.Buffer)
	goModT.Execute(goModBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/go.mod", directoryName), goModBuffer.String())

	util.CreateFile(fmt.Sprintf("%v/adventofcode.go", directoryName), golang.AdventOfCodeTemplate)

	fmt.Println("Created new project for the advent of code", year)
}
