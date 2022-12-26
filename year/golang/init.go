package golang

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/gossie/aoc-generator/templates/golang"
	"github.com/gossie/aoc-generator/util"
)

type fileData struct {
	GithubUser string
	Year       int
}

func InitializeYear(directoryName string, year int, githubUser string) {
	fileData := fileData{githubUser, year}

	goModT, err := template.New("goMod").Parse(golang.GoModTemplate)
	if err != nil {
		panic("template could not be parsed")
	}

	goModBuffer := new(bytes.Buffer)
	goModT.Execute(goModBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/go.mod", directoryName), goModBuffer.String())
	util.CreateFile(fmt.Sprintf("%v/adventofcode.go", directoryName), golang.AdventOfCodeTemplate)
}
