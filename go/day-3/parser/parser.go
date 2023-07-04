package parser

import (
	"os"
	"strings"
)

type Backpack [2]string
type BackpackList []Backpack

func Parse(path string) BackpackList {
	var file, err = os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	rawContent := string(file)
	lines := strings.Split(rawContent, "\n")
	result := make([]Backpack, 0, len(lines))

	for _, line := range lines {
		halfIndex := len(line) / 2
		first := line[:halfIndex]
		second := line[halfIndex:]
		backpack := Backpack{first, second}
		result = append(result, backpack)
	}

	return result
}
