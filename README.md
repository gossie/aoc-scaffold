# aoc-generator
A simple CLI that lets you create scaffolds for advent of code projects and single days.

## Installation

If you habe go installed on your system you could clone the GitHub repository, adjust the path in the install.sh and then execute the install.sh.

## Usage

### aoc generate

To create a new aoc project use the `aoc generate` command and pass the following arguments:
* **language**: The prgramming language you want to use (currently go and java are supported)
* **year**: The year your are implementing the tasks for
* **githubUser**: This is used for the module in the generated go.mod file or the groupId in the pom.xml

This command will create a new directory, that contains your project.

### aoc create

To create the scaffold for a new day run the `aoc create` command inside the directory, that was create by the `aoc generate` command, and pass the following arguments:
* **day**: The day you want to implement

This command will create a new subfolder for the day, that contains a source file with functions for part one and part two, a test file with unit tests for part one and part two and two text files for the puzzle input and the example input. Also, this command will alter the main source file and insert calls to the funtions for part one and two and print the time that they took to run.
