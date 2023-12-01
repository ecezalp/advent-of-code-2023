import * as fs from "fs";

export { Run, Day };

const Run = (day: Day) => {
  const args = process.argv.slice(2);
  day.RunWithInput1(loadFileIntoArray(args[0]));
  day.RunWithInput2(loadFileIntoArray(args[1]));
};

interface Day {
  RunWithInput1: (input: string[]) => void;
  RunWithInput2: (input: string[]) => void;
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
