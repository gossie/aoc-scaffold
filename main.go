package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gossie/aoc-generator/day"
	"github.com/gossie/aoc-generator/year"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	aocYear := generateCmd.Int("year", time.Now().Year(), "The year the project is used for")
	language := generateCmd.String("language", "", "The programming language you are using (currently supported is Go)")
	githubUser := generateCmd.String("githubUser", "", "Your GitHub username (only used for the go module)")

	createDayCmd := flag.NewFlagSet("create", flag.ExitOnError)
	aocDay := createDayCmd.Int("day", time.Now().Day(), "The day you want to implement the task for")

	command := os.Args[1]
	switch command {
	case "generate":
		handleGenerate(generateCmd, aocYear, language, githubUser)
	case "create":
		handleCreateDay(createDayCmd, aocDay)
	case "--help":
		printHelp([]*flag.FlagSet{generateCmd, createDayCmd})
	default:
		printHelp([]*flag.FlagSet{generateCmd, createDayCmd})
	}
}

func handleGenerate(generateCmd *flag.FlagSet, aocYear *int, language, githubUser *string) {
	generateCmd.Parse(os.Args[2:])
	if *language == "go" && *githubUser == "" {
		fmt.Println("I need the GitHub username to create a go module")
		os.Exit(1)
	}
	year.InitializeYear(*aocYear, *language, *githubUser)
}

func handleCreateDay(createDayCmd *flag.FlagSet, aocDay *int) {
	createDayCmd.Parse(os.Args[2:])
	day.CreateDay(*aocDay)
}

func printHelp(commands []*flag.FlagSet) {
	fmt.Println("Use aoc to generate scaffolds for your advent of code go project.")
	for _, cmd := range commands {
		fmt.Println(cmd.Name())
		cmd.PrintDefaults()
	}
}
