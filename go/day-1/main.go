package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := getFilePath()
	file := loadFile(filepath)
	calories := parse(file)
	maxCalories := findBiggestCalories(calories)
	fmt.Printf("Maximum calories: %d\n", maxCalories)
}

func findBiggestCalories(calories [][]int) int {
	var max int

	for _, group := range calories {
		var sum int

		for _, item := range group {
			sum += item
		}

		if max < sum {
			max = sum
		}
	}

	return max
}

func parse(content []byte) [][]int {
	rawContent := string(content)
	words := strings.Split(rawContent, "\n")
	result := make([][]int, 0, len(words)/5)
	temporaryGroup := make([]int, 0, 100)

	for _, word := range words {
		if word == "" {
			result = append(result, temporaryGroup)
			temporaryGroup = make([]int, 0, 100)
			continue
		}

		calories, err := strconv.Atoi(word)

		if err != nil {
			panic(err)
		}

		temporaryGroup = append(temporaryGroup, calories)
	}

	return result
}

func getFilePath() string {
	if len(os.Args) < 2 {
		panic("Filepath not specified in vArgs")
	}

	return os.Args[1]
}

func loadFile(filepath string) []byte {
	content, err := os.ReadFile(filepath)

	if err != nil {
		message := fmt.Sprintf("There's no such file with path %v", filepath)
		panic(message)
	}

	return content
}
