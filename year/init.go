package year

import (
	"fmt"
	"os"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/year/golang"
)

var initializer = map[string]func(string, int, string, string){
	"go": golang.InitializeYear,
}

func InitializeYear(year int, language, githubUser string) {
	directoryName := fmt.Sprintf("advent-of-code-%d", year)
	err := os.Mkdir(directoryName, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + directoryName)
	}

	config.WriteConfig(map[string]string{"year": fmt.Sprintf("%d", year), "language": language, "githubUser": githubUser})

	initializer[language](directoryName, year, language, githubUser)

	fmt.Println("Created new project for the advent of code", year)
}
