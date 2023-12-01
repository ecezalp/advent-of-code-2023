package days

import (
	"fmt"
	"os"
)

func (d Day1Module) RunWithInput1() {
	fmt.Printf("Running day 1 function with input 1: %v\n", d.input1)
}

func (d Day1Module) RunWithInput2() {
	if len(d.input2) == 0 {
		fmt.Println("Input 2 not yet unlocked")
		os.Exit(0)
	}
	fmt.Printf("Running day 1 function with input 2: %v\n", d.input2)
}

func InitDay1Module(input1 []string, input2 []string) Day1Module {
	return Day1Module{
		input1: input1,
		input2: input2,
	}
}

type Day1Module struct {
	input1 []string
	input2 []string
}
