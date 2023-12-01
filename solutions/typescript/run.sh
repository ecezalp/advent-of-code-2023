#!/bin/bash

# Validate that a day number is provided
if [ "$#" -ne 1 ]; then
  echo "Please provide the day number as an argument: $0 <day_number>"
  exit 1
fi

# Get the day number from the command-line argument
DAY_NUMBER="$1"

# Get the current directory of the script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
INPUT_DIR="$(cd "../../input" && pwd )"

# Run ts-node with the specified day file
npx ts-node --project "$SCRIPT_DIR/tsconfig.json" "$SCRIPT_DIR/days/$DAY_NUMBER.ts" "$INPUT_DIR/day${DAY_NUMBER}_input1.txt" "$INPUT_DIR/day${DAY_NUMBER}_input2.txt"