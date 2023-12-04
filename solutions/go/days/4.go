package days

import (
	"fmt"
	"math"
	"strings"
)

type Day4 struct{}

func (d Day4) Part1(input []string) string {
	var total float64
	for _, line := range input {
		parts := strings.Split(line[8:len(line)], " | ")
		winning := make(map[string]bool)
		winningStrings := strings.Split(parts[0], " ")
		for _, winningString := range winningStrings {
			if winningString != "" {
				winning[winningString] = true
			}
		}
		testStrings := strings.Split(parts[1], " ")
		fmt.Println(winningStrings, testStrings)
		var matchCount float64
		for _, testString := range testStrings {
			if winning[testString] {
				matchCount++
			}
		}
		total += math.Floor(math.Exp2(matchCount - 1))
	}
	return fmt.Sprintf("%v", total)
}

func (d Day4) Part2(input []string) string {
	matchMap := make(map[int]int)
	for i, line := range input {
		parts := strings.Split(line[8:len(line)], " | ")
		winning := make(map[string]bool)
		winningStrings := strings.Split(parts[0], " ")
		for _, winningString := range winningStrings {
			if winningString != "" {
				winning[winningString] = true
			}
		}
		testStrings := strings.Split(parts[1], " ")
		fmt.Println(winningStrings, testStrings)
		var matchCount int
		for _, testString := range testStrings {
			if winning[testString] {
				matchCount++
			}
		}
		matchMap[i+1] = matchCount
	}

	process := []int{}
	for i := 1; i < len(matchMap)+1; i++ {
		process = append(process, i)
	}

	for i := 0; i < len(process); i++ {
		if (matchMap[process[i]]) > 0 {
			for j := 1; j < matchMap[process[i]]+1; j++ {
				if (process[i] + j) < len(input)+1 {
					process = append(process, process[i]+j)
				}
			}
		}
	}
	return fmt.Sprintf("%v", len(process))
}

func NewDay4() Day4 {
	return Day4{}
}
