package year

import (
	"fmt"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/util"
	"github.com/gossie/aoc-generator/year/golang"
	"github.com/gossie/aoc-generator/year/java"
)

var initializer = map[string]func(string, int, string){
	"go":   golang.InitializeYear,
	"java": java.InitializeYear,
}

func InitializeYear(year int, language, githubUser string) {
	directoryName := fmt.Sprintf("advent-of-code-%d", year)
	util.CreateDirectory(directoryName)

	config.WriteConfig(map[string]string{"year": fmt.Sprintf("%d", year), "language": language, "githubUser": githubUser})

	initializer[language](directoryName, year, githubUser)

	fmt.Println("Created new project for the advent of code", year)
}
