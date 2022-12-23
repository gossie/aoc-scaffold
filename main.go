package main

import (
	"os"

	"github.com/gossie/aoc-generator/day"
	"github.com/gossie/aoc-generator/year"
)

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		if len(os.Args) == 4 {
			year.InitializeYear(os.Args[2], os.Args[3])
		} else {
			println("I will help you later, I promise")
		}
	case "create":
		if len(os.Args) == 3 {
			day.CreateDay(os.Args[2])
		} else {
			println("I will help you later, I promise")
		}
	case "--help":
		printHelp()
	default:
		printHelp()
	}
}

func printHelp() {
	println("Use aoc to generate scaffolds for your advent of code go project.")
	println()
	println("\tinit <year> <githubUser>\tCreates a new advent of code project.")
	println("\t\t\t\t\tasd")
	println("\tcreate <day>\t\t\tCreates a new day within your advent of code project.")
}
