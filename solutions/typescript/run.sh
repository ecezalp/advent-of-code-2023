#!/bin/bash

# Function to display usage information
usage() {
  echo "Usage: $0 --day <day_number> --part <1|2> [--test]"
  exit 1
}

# Validate that at least one argument is provided
if [ "$#" -lt 1 ]; then
  usage
fi

# Initialize variables
DAY_NUMBER=""
PART=""
TEST_FLAG=""

# Process command-line arguments
while [ "$#" -gt 0 ]; do
  case "$1" in
    --day)
      shift
      DAY_NUMBER="$1"
      ;;
    --part)
      shift
      PART="$1"
      ;;
    --test)
      TEST_FLAG="--test"
      ;;
    *)
      usage
      ;;
  esac
  shift
done

# Validate mandatory arguments
if [ -z "$DAY_NUMBER" ] || [ -z "$PART" ]; then
  usage
fi

# Validate --part argument
if [ "$PART" -ne 1 ] && [ "$PART" -ne 2 ]; then
  echo "--part must be either 1 or 2."
  usage
fi

# Get input strings
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
INPUT_DIR="$(cd "../../input" && pwd )"

# Set input and answer file paths based on the presence of the --test flag
if [ -n "$TEST_FLAG" ]; then
  INPUT_FILE="$INPUT_DIR/day${DAY_NUMBER}_test${PART}.txt"
  ANSWER_FILE="$INPUT_DIR/day${DAY_NUMBER}_test${PART}_answer.txt"
else
  INPUT_FILE="$INPUT_DIR/day${DAY_NUMBER}_input${PART}.txt"
  ANSWER_FILE=""  # No answer file for non-test runs
fi

# Run ts-node with the specified day file
npx ts-node --project "$SCRIPT_DIR/tsconfig.json" "$SCRIPT_DIR/days/$DAY_NUMBER.ts" "$PART" "$INPUT_FILE" "$ANSWER_FILE"