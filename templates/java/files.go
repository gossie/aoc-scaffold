package java

import (
	_ "embed"
)

//go:embed pom.xml.template
var PomTemplate string

//go:embed AdventOfCode.java.template
var AdventOfCodeTemplate string

//go:embed SingleDay.part.template
var SingleDayInAdventOfCodeTemplate string

//go:embed Day.java.template
var DayTemplate string

//go:embed DayTest.java.template
var DayTestTemplate string
