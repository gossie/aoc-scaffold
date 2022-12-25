package java

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gossie/aoc-generator/templates/java"
	"github.com/gossie/aoc-generator/util"
)

type fileData struct {
	GithubUser string
	Year       int
}

func InitializeYear(directoryName string, year int, githubUser string) {
	fileData := fileData{githubUser, year}

	pomT, err := template.New("pom").Parse(java.PomTemplate)
	if err != nil {
		panic("template could not be parsed")
	}

	pomBuffer := new(bytes.Buffer)
	pomT.Execute(pomBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/pom.xml", directoryName), pomBuffer.String())
	//util.CreateFile(fmt.Sprintf("%v/adventofcode.go", directoryName), golang.AdventOfCodeTemplate)
}
