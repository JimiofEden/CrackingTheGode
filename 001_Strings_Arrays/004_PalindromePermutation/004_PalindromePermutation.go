package ch01

import (
	"fmt"
	"strings"
	"sort"
)
/**
 * Will determine if one string is a permutation of another
 */

func palindromePermutationNLogN(inputString string) bool {
	if inputString == "" {
		return false
	}
	inputString = SortString(strings.ToLower(inputString))
	fmt.Println(inputString)

	storedLetter := inputString[0]
	counter := 1
	singleOdd := false

	for i := 1; i < len(inputString); i++ {
		if inputString[i] == ' ' {
			continue
		}
		if storedLetter == ' ' {
			storedLetter = inputString[i]
			counter = 1
			continue
		}
		if inputString[i] != storedLetter {
			if counter%2 != 0 {
				if singleOdd { 
					return false
				}
				singleOdd = true
			}
			storedLetter = inputString[i]
			counter = 1
			continue
		}
		counter++
	}

	return true
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func palindromePermutationN(inputString string) bool {
	if inputString == "" {
		return false
	}
	inputString = strings.ToLower(inputString)

	charMap := make(map[rune]int)

	for _, l := range inputString {
		if l != ' ' {
			charMap[rune(l)]++
		}
	}

	singleOdd := false
	for _, v := range charMap {
		if v%2 != 0 {
			if singleOdd {
				return false
			}
			singleOdd = true
		}
	}

	return true
}

