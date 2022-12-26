package java

import (
	"bytes"
	"fmt"
	"text/template"

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

	util.CreateDirectory(fmt.Sprintf("%v/src", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/main", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/main/resources", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/main/resources/adventofcode", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/main/java", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/main/java/adventofcode", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/test", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/test/resources", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/test/resources/adventofcode", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/test/java", directoryName))
	util.CreateDirectory(fmt.Sprintf("%v/src/test/java/adventofcode", directoryName))

	pomBuffer := new(bytes.Buffer)
	pomT.Execute(pomBuffer, fileData)
	util.CreateFile(fmt.Sprintf("%v/pom.xml", directoryName), pomBuffer.String())
	util.CreateFile(fmt.Sprintf("%v/src/main/java/adventofcode/AdventOfCode.java", directoryName), java.AdventOfCodeTemplate)
}
