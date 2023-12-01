#!/bin/bash

# Validate that a day number is provided
if [ "$#" -ne 1 ]; then
  echo "Please provide the day number as an argument: $0 <day_number>"
  exit 1
fi

# Get the day number from the command-line argument
DAY_NUMBER="$1"

# Get input strings
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
INPUT_DIR="$(cd "../../input" && pwd )"

# Run main.go with the specified day file and input files
go run main.go "$DAY_NUMBER" "$INPUT_DIR/day${DAY_NUMBER}_input1.txt" "$INPUT_DIR/day${DAY_NUMBER}_input2.txt"