package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Top [3]int

func main() {
    filepath := getFilePath()
    file := loadFile(filepath)
    calories := parse(file)
    maxCalories := findBiggestCalories(calories)
    top3 := findTop(calories)
    fmt.Printf("Maximum calories: %d\n", maxCalories)
    fmt.Printf("Sum of top 3 elves calories is: %d\n", findTopSum(top3))
}

func findTopSum(top Top) int {
    var sum int

    for _, item := range top {
        sum += item
    }

    return sum
}

func findTop(calories [][]int) Top {
    var top [3]int

    for _, group := range calories {
        var sum int

        for _, item := range group {
            sum += item
        }

        // Check top iterms
        for idx, topItem := range top {
            if topItem < sum {
                top[idx] = sum
                break
            }
        }
    }

    fmt.Println(top)
    return top
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

    if len(temporaryGroup) > 0 {
        result = append(result, temporaryGroup)
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
