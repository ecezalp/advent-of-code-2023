package main

import (
	"bufio"
	. "ece-advent-2023/days"
	"fmt"
	"os"
)

func main() {
	dayNumber := os.Args[1]
	part := os.Args[2]
	input := createSliceFromInput(os.Args[3])
	output := ""
	if os.Args[4] != "" {
		output = createSliceFromInput(os.Args[4])[0]
	}

	day := getDay(dayNumber)
	var run func(input []string) string
	if part == "1" {
		run = day.Part1
	} else {
		run = day.Part2
	}

	result := run(input)
	if output != "" {
		if result == output {
			fmt.Println("Success")
		} else {
			fmt.Println(fmt.Sprintf("Fail: Expected %s, Got %s", output, result))
		}
	} else {
		fmt.Println(result)
	}
}

func getDay(dayNumber string) Day {
	var day Day
	switch dayNumber {
	case "1":
		day = NewDay1()
	case "2":
		day = NewDay2()
	case "3":
		day = NewDay3()
	case "4":
		day = NewDay4()
	case "5":
		day = NewDay5()
	case "6":
		day = NewDay6()
	case "7":
		day = NewDay7()
	case "8":
		day = NewDay8()
	case "10":
		day = NewDay10()
	default:
		day = NewDay1()
	}
	return day
}

func createSliceFromInput(filePath string) (response []string) {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
		return
	}
	defer file.Close()

	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		response = append(response, line)
	}

	// Handle error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return response
}

type Day interface {
	Part1(input []string) string
	Part2(input []string) string
}
