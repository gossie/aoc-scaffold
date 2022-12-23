package golang

import (
	_ "embed"
)

//go:embed go.mod.template
var GoModTemplate string

//go:embed adventofcode.go.template
var AdventOfCodeTemplate string

//go:embed singleDay.part.template
var SingleDayInAdventOfCodeTemplate string

//go:embed day.go.template
var DayTemplate string

//go:embed day_test.go.template
var DayTestTemplate string
