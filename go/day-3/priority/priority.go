package priority

import (
	"fmt"
	"strings"
)

// Magical numbers ✨
// If we substract these number, from their ASCII code we should get
// a-z -> 1..=26
// A-Z -> 27..=52
const uppercaseModifier = 38
const lowercaseModifier = 96

func GetPriority(character rune) int {
	var modifier int

	// Check if letter is lowercase, uppercase characters has higher ascii code

	if rune(strings.ToLower(string(character))[0]) == character {
		modifier = lowercaseModifier
		fmt.Println("Lowercase character", string(character))
	} else {
		modifier = uppercaseModifier
		fmt.Println("Uppercase character", string(character))
	}

	return int(character) - modifier
}
