package main

import (
	"bufio"
	"ece-advent-2023/days"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		os.Exit(1)
	}
	dayRunner := loadDay(os.Args[1], createSliceFromInput(os.Args[2]), createSliceFromInput(os.Args[3]))
	dayRunner.RunWithInput1()
	dayRunner.RunWithInput2()
}

func loadDay(dayNumber string, input1 []string, input2 []string) DayRunner {
	switch dayNumber {
	case "1":
		return days.InitDay1Module(input1, input2)
	default:
		return days.InitDay1Module(input1, input2)
	}
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

type DayRunner interface {
	RunWithInput1()
	RunWithInput2()
}
