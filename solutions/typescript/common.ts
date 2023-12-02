import * as fs from "fs";

export { Run, Day };

const Run = (day: Day) => {
  const args = process.argv.slice(2);
  const part = args[0];
  const input = loadFileIntoArray(args[1]);
  let output = "";
  if (args[2] !== "") {
    output = loadFileIntoArray(args[2])[0];
  }

  let run: (input: string[]) => string;
  if (part == "1") {
    run = day.Part1;
  } else {
    run = day.Part2;
  }

  const result = run(input);
  if (output != "") {
    if (result == output) {
      console.log("Success");
    } else {
      console.log("Fail: expected " + output + ", found " + result);
    }
  } else {
    console.log(result);
  }
};

interface Day {
  Part1: (input: string[]) => string;
  Part2: (input: string[]) => string;
}

function loadFileIntoArray(filePath: string): string[] {
  try {
    const fileContent = fs.readFileSync(filePath, "utf-8");
    const lines = fileContent.split("\n");
    return lines;
  } catch (error) {
    console.error(`Error reading file ${filePath}:`, error);
    return [];
  }
}
