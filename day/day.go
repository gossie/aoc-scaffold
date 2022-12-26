package day

import (
	"github.com/gossie/aoc-generator/config"
	"github.com/gossie/aoc-generator/day/golang"
	"github.com/gossie/aoc-generator/day/java"
)

var initializer = map[string]func(int){
	"go":   golang.CreateDay,
	"java": java.CreateDay,
}

func CreateDay(aocDay int) {
	language := config.GetPropertyValue("language")
	initializer[language](aocDay)
}
