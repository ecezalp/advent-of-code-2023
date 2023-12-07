package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day5 struct{}

func (d Day5) Part1(input []string) string {
	values := make([][]int, 8)
	lineBreakCount := 0
	for i := 0; i < len(input); i++ {
		existingValues := values[lineBreakCount]
		numberStrings := strings.Split(input[i], " ")
		for _, numberString := range numberStrings {
			if numberString != "" {
				number, _ := strconv.Atoi(numberString)
				existingValues = append(existingValues, number)
			}
		}
		values[lineBreakCount] = existingValues
		if input[i] == "" {
			lineBreakCount++
			i = i + 1
		}
	}

	transformed := []int{}
	transform := func(valueSlice []int, currentVal int) (newVal int) {
		for i := 0; i < len(valueSlice); i = i + 3 {
			source := valueSlice[i+1]
			rangeLen := valueSlice[i+2]

			if currentVal >= source && currentVal <= source+rangeLen {
				diff := source - valueSlice[i]
				newVal = currentVal - diff
				return newVal
			}
		}
		return currentVal
	}

	for i := 1; i < len(values[0]); i++ {
		value := values[0][i]
		for j := 1; j < 8; j++ {
			value = transform(values[j], value)
			fmt.Println(value)
		}
		transformed = append(transformed, value)
	}

	sort.Ints(transformed)
	return fmt.Sprintf("%v", transformed[0])
}

func (d Day5) Part2(input []string) string {
	values := make([][]int, 8)
	lineBreakCount := 0
	for i := 0; i < len(input); i++ {
		existingValues := values[lineBreakCount]
		numberStrings := strings.Split(input[i], " ")
		for _, numberString := range numberStrings {
			if numberString != "" {
				number, _ := strconv.Atoi(numberString)
				existingValues = append(existingValues, number)
			}
		}
		values[lineBreakCount] = existingValues
		if input[i] == "" {
			lineBreakCount++
			i = i + 1
		}
	}
	values[0] = values[0][1:]

	findLowestValueRange := func(valueSlice []int) (source int, dest int) {
		lowestDest := 0
		lowestSource := 0
		for i := 0; i < len(valueSlice); i = i + 3 {
			if i == 0 {
				lowestDest = valueSlice[i+1]
			}
			if valueSlice[i+1] < lowestDest {
				lowestDest = valueSlice[i+1]
			}
		}
		return lowestSource, lowestDest
	}

	findCorrespondingRange := func(valueSlice []int, prevStart int,
		prevEnd int) (nextStart int, nextEnd int) {
		closestDestIndex := -1
		closestDestDiff := -1
		for i := 0; i < len(valueSlice); i = i + 3 {
			if valueSlice[i]-prevStart < closestDestDiff || closestDestIndex == -1 {
				closestDestDiff = valueSlice[i] - prevStart
				closestDestIndex = i
			}
		}

		nextStart = valueSlice[closestDestIndex+1]
		nextEnd = nextStart
		diff := math.Min(float64(prevEnd-prevStart), float64(valueSlice[closestDestIndex+2]))

		for i := 0; i < int(diff); i++ {
			nextEnd++
		}

		return nextStart, nextEnd
	}

	start, end := findLowestValueRange(values[7])
	source, dest := findCorrespondingRange(values[6], start, end)
	source, dest = findCorrespondingRange(values[5], source, dest)
	source, dest = findCorrespondingRange(values[4], source, dest)
	source, dest = findCorrespondingRange(values[3], source, dest)
	source, dest = findCorrespondingRange(values[2], source, dest)
	source, dest = findCorrespondingRange(values[1], source, dest)

	// 1520731987, 1594903067

	return fmt.Sprintf("%v, %v", Day5.Part1())
}

func NewDay5() Day5 {
	return Day5{}
}
