package main

import (
	"fmt"
	"main/backpackSort"
	"main/cli"
	"main/parser"
	"main/priority"
	"sort"
)

func transformBackpackCharacters(backpack parser.Backpack) (result backpackSort.RuneBackpack) {
	for partIdx, backpackPart := range backpack {
		runeBackpackPart := make([]rune, 0, len(backpackPart))

		for _, character := range backpackPart {
			runeBackpackPart = append(runeBackpackPart, character)
		}

		result[partIdx] = runeBackpackPart
	}

	return
}

func findCommon(backpacks parser.BackpackList) []rune {
	var commonCharacters []rune

	for _, backpack := range backpacks {

		// Get runes from characters in backpack
		runeBackpack := transformBackpackCharacters(backpack)

		// Sort backpack
		for _, part := range runeBackpack {
			sort.Sort(&part)
		}

		// Select primary backpack
		var primaryBackpackPart backpackSort.RuneBackpackPart
		var secondaryBackpackPart backpackSort.RuneBackpackPart

		if len(backpack[0]) < len(backpack[1]) {
			primaryBackpackPart = runeBackpack[0]
			secondaryBackpackPart = runeBackpack[1]
		} else {
			primaryBackpackPart = runeBackpack[1]
			secondaryBackpackPart = runeBackpack[0]
		}

		for _, character := range primaryBackpackPart {
			foundIdx, found := sort.Find(secondaryBackpackPart.Len(), func(i int) int {
				return int(character - secondaryBackpackPart[i])
			})

			if found {
				commonCharacters = append(commonCharacters, secondaryBackpackPart[foundIdx])
				break
			}
		}
	}

	return commonCharacters
}

func main() {
	path := cli.GetPath()
	backpacks := parser.Parse(path)
	commonRunes := findCommon(backpacks)
	fmt.Println(commonRunes)

	var sum int
	for _, commonRune := range commonRunes {
		sum += priority.GetPriority(commonRune)
	}

	fmt.Println(sum)
}
