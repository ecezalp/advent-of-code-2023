package days

import "strings"

type Day8 struct{}

func (d Day8) Part1(input []string) string {
	var instructions []string
	right := make(map[string]string)
	left := make(map[string]string)

	for i, line := range input {
		if i == 0 {
			instructions = strings.Split(line)
		}
		if i > 1 {

		}
	}
	return ""
}

func (d Day8) Part2(input []string) string {
	return ""
}

func NewDay8() Day8 {
	return Day8{}
}
