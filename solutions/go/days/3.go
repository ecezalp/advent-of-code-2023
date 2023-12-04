package days

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"unicode"
)

type Day3 struct{}

func (d Day3) Part1(input []string) string {
	total := 0
	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if unicode.IsNumber(rune(line[j])) {
				numberString := ""
			kLoop:
				for k := j; k < len(line); k++ {
					if !unicode.IsNumber(rune(line[k])) {
						break kLoop
					} else {
						numberString += string(line[k])
					}
				}

				addNumber := false
				topI := i - 1
				bottomI := i + 1
				leftJ := j - 1
				rightJ := j + len(numberString)

				pattern := `[^0-9\.]`
				regex := regexp.MustCompile(pattern)

				// left
				if leftJ > -1 && regex.MatchString(string(line[leftJ])) {
					addNumber = true
				}

				// right
				if rightJ < len(line) {
					if regex.MatchString(string(line[rightJ])) {
						addNumber = true
					}
				}

				// top slice
				if topI > -1 {
					sliceStart := leftJ
					if sliceStart < 0 {
						sliceStart = 0
					}
					sliceEnd := rightJ + 1
					if sliceEnd > len(line) {
						sliceEnd = rightJ
					}
					if regex.MatchString(input[topI][sliceStart:sliceEnd]) {
						addNumber = true
					}
				}

				// bottom slice
				if bottomI < len(input) {
					sliceStart := leftJ
					if sliceStart < 0 {
						sliceStart = 0
					}
					sliceEnd := rightJ + 1
					if sliceEnd > len(line) {
						sliceEnd = rightJ
					}
					if regex.MatchString(input[bottomI][sliceStart:sliceEnd]) {
						addNumber = true
					}
				}

				if addNumber == true {
					num, _ := strconv.Atoi(numberString)
					total += num
				}
				j = rightJ
			}
		}
	}
	return fmt.Sprintf("%v", total)
}

func (d Day3) Part2(input []string) string {
	type Coordinate struct {
		X int
		Y int
	}
	numberCoordinates := make(map[Coordinate]string)
	valueMap := make(map[string]int)

	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if unicode.IsNumber(rune(line[j])) {
				id := uuid.New()
				numberString := ""
			kLoop:
				for k := j; k < len(line); k++ {
					if !unicode.IsNumber(rune(line[k])) {
						break kLoop
					} else {
						numberCoordinates[Coordinate{
							X: i,
							Y: k,
						}] = id.String()
						numberString += string(line[k])
						j = k
					}
				}
				val, _ := strconv.Atoi(numberString)
				valueMap[id.String()] = val
			}
		}
	}

	total := 0
	starCoordinates := []Coordinate{}

	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if string(line[j]) == "*" {
				starCoordinates = append(starCoordinates, Coordinate{X: i, Y: j})
			}
		}
	}

	for _, star := range starCoordinates {
		uniqueNeighbors := make(map[string]bool)
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X - 1,
			Y: star.Y - 1,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X - 1,
			Y: star.Y,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X - 1,
			Y: star.Y + 1,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X,
			Y: star.Y - 1,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X,
			Y: star.Y + 1,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X + 1,
			Y: star.Y - 1,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X + 1,
			Y: star.Y,
		}]] = true
		uniqueNeighbors[numberCoordinates[Coordinate{
			X: star.X + 1,
			Y: star.Y + 1,
		}]] = true

		delete(uniqueNeighbors, "")
		if len(uniqueNeighbors) == 2 {
			fmt.Printf("%v", star)
			ratio := 1
			for uuidVal := range uniqueNeighbors {
				ratio *= valueMap[uuidVal]
			}
			total += ratio
		}
	}

	return fmt.Sprintf("%v", total)
}

func NewDay3() Day3 {
	return Day3{}
}
