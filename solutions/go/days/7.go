package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct{}

func (d Day7) Part1(input []string) string {
	scores := make(map[string]int)
	var keys []string

	for _, line := range input {
		parts := strings.Split(line, " ")
		score, _ := strconv.Atoi(parts[1])
		scores[parts[0]] = score
		keys = append(keys, parts[0])
	}

	getStrength := func(card string) int {
		chars := strings.Split(card, "")
		charCount := make(map[string]int)
		for _, char := range chars {
			charCount[char] = charCount[char] + 1
		}
		if len(charCount) == 1 {
			// 5 of a kind
			return 10
		}
		if len(charCount) == 2 {
			for _, v := range charCount {
				if v == 4 {
					// 4 of a kind
					return 9
				}
			}
			// full house
			return 8
		}
		if len(charCount) == 3 {
			for _, v := range charCount {
				if v == 3 {
					// 3 of a kind
					return 7
				}
			}
			// 2 pairs
			return 6
		}
		if len(charCount) == 4 {
			// one pair
			return 5
		}
		// high card
		return 4
	}

	charValues := make(map[string]int)
	charValues["2"] = 0
	charValues["3"] = 1
	charValues["4"] = 2
	charValues["5"] = 3
	charValues["6"] = 4
	charValues["7"] = 5
	charValues["8"] = 6
	charValues["9"] = 7
	charValues["T"] = 8
	charValues["J"] = 9
	charValues["Q"] = 10
	charValues["K"] = 11
	charValues["A"] = 12

	sort.Slice(keys, func(i, j int) bool {
		iStrength := getStrength(keys[i])
		jStrength := getStrength(keys[j])

		if iStrength < jStrength {
			return true
		}
		if jStrength < iStrength {
			return false
		}

		for k := 0; k < 5; k++ {
			if charValues[string(keys[i][k])] < charValues[string(keys[j][k])] {
				return true
			}
			if charValues[string(keys[j][k])] < charValues[string(keys[i][k])] {
				return false
			}
		}
		return true
	})

	fmt.Println(keys)
	total := 0
	for i := 0; i < len(keys); i++ {
		val := scores[keys[i]] * (i + 1)
		fmt.Println(total, val)
		total = total + val
	}

	return fmt.Sprintf("%d", total)
}

func (d Day7) Part2(input []string) string {
	scores := make(map[string]int)
	var keys []string

	for _, line := range input {
		parts := strings.Split(line, " ")
		score, _ := strconv.Atoi(parts[1])
		scores[parts[0]] = score
		keys = append(keys, parts[0])
	}

	getStrength := func(card string) int {
		chars := strings.Split(card, "")
		charCount := make(map[string]int)
		jCount := 0
		for _, char := range chars {
			if char == "J" {
				jCount++
			} else {
				charCount[char] = charCount[char] + 1
			}
		}

		if jCount > 0 {
			highestCount := 0
			highestChar := ""
			for k, v := range charCount {
				if charCount[k] > highestCount {
					highestChar = k
					highestCount = v
				}
			}
			charCount[highestChar] = highestCount + jCount
		}

		if len(charCount) == 1 {
			// 5 of a kind
			return 10
		}
		if len(charCount) == 2 {
			for _, v := range charCount {
				if v == 4 {
					// 4 of a kind
					return 9
				}
			}
			// full house
			return 8
		}
		if len(charCount) == 3 {
			for _, v := range charCount {
				if v == 3 {
					// 3 of a kind
					return 7
				}
			}
			// 2 pairs
			return 6
		}
		if len(charCount) == 4 {
			// one pair
			return 5
		}
		// high card
		return 4
	}

	charValues := make(map[string]int)
	charValues["2"] = 1
	charValues["3"] = 2
	charValues["4"] = 3
	charValues["5"] = 4
	charValues["6"] = 5
	charValues["7"] = 6
	charValues["8"] = 7
	charValues["9"] = 8
	charValues["T"] = 9
	charValues["J"] = 0
	charValues["Q"] = 10
	charValues["K"] = 11
	charValues["A"] = 12

	sort.Slice(keys, func(i, j int) bool {
		iStrength := getStrength(keys[i])
		jStrength := getStrength(keys[j])

		if iStrength < jStrength {
			return true
		}
		if jStrength < iStrength {
			return false
		}

		for k := 0; k < 5; k++ {
			if charValues[string(keys[i][k])] < charValues[string(keys[j][k])] {
				return true
			}
			if charValues[string(keys[j][k])] < charValues[string(keys[i][k])] {
				return false
			}
		}
		return true
	})

	fmt.Println(keys)
	total := 0
	for i := 0; i < len(keys); i++ {
		val := scores[keys[i]] * (i + 1)
		fmt.Println(total, val)
		total = total + val
	}

	return fmt.Sprintf("%d", total)
}

func NewDay7() Day7 {
	return Day7{}
}
