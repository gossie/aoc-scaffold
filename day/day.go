package day

import (
	"os"

	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/day/golang"
)

var initializer = map[string]func(string){
	"go": golang.CreateDay,
}

func CreateDay(name string) {
	err := os.Mkdir(name, os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory " + name)
	}

	language := config.GetPropertyValue("language")
	initializer[language](name)
}
