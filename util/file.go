package util

import (
	"os"
)

func CreateFile(name, content string) {
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
