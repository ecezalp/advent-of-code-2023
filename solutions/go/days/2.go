package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Part1(input []string) string {
	total := 0
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

gameLoop:
	for _, line := range input {
		splitLine := strings.Split(line, ":")
		id, _ := strconv.Atoi(splitLine[0][5:])

		gameString := strings.Replace(splitLine[1], " ", "", -1)
		roundStrings := strings.Split(gameString, ";")
		for _, roundString := range roundStrings {
			colorStrings := strings.Split(roundString, ",")
			for _, colorString := range colorStrings {
				if colorString[len(colorString)-3:] == "red" {
					redCount, _ := strconv.Atoi(colorString[:len(colorString)-3])
					if redCount > redLimit {
						continue gameLoop
					}
				}
				if len(colorString) > 4 && colorString[len(colorString)-4:] == "blue" {
					blueCount, _ := strconv.Atoi(colorString[:len(colorString)-4])
					if blueCount > blueLimit {
						continue gameLoop
					}
				}
				if len(colorString) > 5 && colorString[len(colorString)-5:] == "green" {
					greenCount, _ := strconv.Atoi(colorString[:len(colorString)-5])
					if greenCount > greenLimit {
						continue gameLoop
					}
				}
			}
		}
		total += id
	}
	return fmt.Sprintf("%d", total)
}

func (d Day2) Part2(input []string) string {
	total := 0
	for _, line := range input {
		redMax := 0
		blueMax := 0
		greenMax := 0

		splitLine := strings.Split(line, ":")
		gameString := strings.Replace(splitLine[1], " ", "", -1)
		roundStrings := strings.Split(gameString, ";")
		for _, roundString := range roundStrings {
			colorStrings := strings.Split(roundString, ",")
			for _, colorString := range colorStrings {
				if colorString[len(colorString)-3:] == "red" {
					redCount, _ := strconv.Atoi(colorString[:len(colorString)-3])
					if redCount > redMax {
						redMax = redCount
					}
				}
				if len(colorString) > 4 && colorString[len(colorString)-4:] == "blue" {
					blueCount, _ := strconv.Atoi(colorString[:len(colorString)-4])
					if blueCount > blueMax {
						blueMax = blueCount
					}
				}
				if len(colorString) > 5 && colorString[len(colorString)-5:] == "green" {
					greenCount, _ := strconv.Atoi(colorString[:len(colorString)-5])
					if greenCount > greenMax {
						greenMax = greenCount
					}
				}
			}
		}
		total += redMax * greenMax * blueMax
	}
	return fmt.Sprintf("%d", total)
}

func NewDay2() Day2 {
	return Day2{}
}
