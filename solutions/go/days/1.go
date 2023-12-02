package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d Day1) Part1(input []string) string {
	isNumber := func(char string) bool {
		if _, err := strconv.Atoi(char); err == nil {
			return true
		}
		return false
	}

	total := 0
	for _, line := range input {
		firstNumber := ""
		lastNumber := ""
		for _, character := range strings.Split(line, "") {
			if isNumber(character) {
				if firstNumber == "" {
					firstNumber = character
				} else {
					lastNumber = character
				}
			}
		}
		if lastNumber == "" {
			lastNumber = firstNumber
		}
		itemAsNumber, _ := strconv.Atoi(firstNumber + lastNumber)
		total += itemAsNumber
	}

	return fmt.Sprintf("%v", total)
}

func (d Day1) Part2(input []string) string {
	replaceThreeLetterNumbers := func(original []string, i int) (updated []string) {
		updated = original
		threeLetterNumbers := map[string]string{
			"one": "1",
			"two": "2",
			"six": "6",
		}
		searched := original[i-2] + original[i-1] + original[i]
		found := threeLetterNumbers[searched]
		if found != "" {
			updated[i-2] = found
		}
		return updated
	}

	replaceFourLetterNumbers := func(original []string, i int) (updated []string) {
		updated = original
		fourLetterNumbers := map[string]string{
			"four": "4",
			"five": "5",
			"nine": "9",
		}
		searched := original[i-3] + original[i-2] + original[i-1] + original[i]
		found := fourLetterNumbers[searched]
		if found != "" {
			updated[i-3] = found
		}
		return updated
	}

	replaceFiveLetterNumbers := func(original []string, i int) (updated []string) {
		updated = original
		fourLetterNumbers := map[string]string{
			"three": "3",
			"seven": "7",
			"eight": "8",
		}
		searched := original[i-4] + original[i-3] + original[i-2] + original[i-1] + original[i]
		found := fourLetterNumbers[searched]
		if found != "" {
			updated[i-4] = found
		}
		return updated
	}

	var updatedInput []string
	for _, inputItem := range input {
		var replacedChars []string
		for i, char := range strings.Split(inputItem, "") {
			replacedChars = append(replacedChars, char)
			if i > 1 {
				replacedChars = replaceThreeLetterNumbers(replacedChars, i)
			}
			if i > 2 {
				replacedChars = replaceFourLetterNumbers(replacedChars, i)
			}
			if i > 3 {
				replacedChars = replaceFiveLetterNumbers(replacedChars, i)
			}
		}
		updatedInput = append(updatedInput, strings.Join(replacedChars, ""))
	}

	return d.Part1(updatedInput)
}

func NewDay1() Day1 {
	return Day1{}
}
