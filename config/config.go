package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteConfig(props map[string]string) {
	err := os.Mkdir(".aoc", os.ModeDir|os.ModePerm)
	if err != nil {
		panic("failed to create directory .aoc")
	}

	content := ""
	for k, v := range props {
		content += fmt.Sprintf("%v=%v\n", k, v)
	}

	createFile(".aoc/config.properties", content)
}

func GetPropertyValue(key string) string {
	return readProperties()[key]
}

func SetPropertyValue(key, value string) {
	props := readProperties()
	props[key] = value
	WriteConfig(props)
}

func createFile(name, content string) {
	testInput, err := os.Create(name)
	if err != nil {
		panic("failed to create file [" + name + "]: " + err.Error())
	}

	_, err = testInput.WriteString(content)
	if err != nil {
		panic("failed to wirte into file [" + name + "]: " + err.Error())
	}
	testInput.Sync()

	defer testInput.Close()
}

func readProperties() map[string]string {
	file, err := os.Open(".aoc/config.properties")
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	properties := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, "=")
		properties[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
	}

	return properties
}
