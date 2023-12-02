import { Run, Day } from "../common";

const Day1: Day = {
  Part1: (input: Array<string>): string => {
    const numbers = ["1", "2", "3", "4", "5", "6", "7", "8", "9"];
    return input
      .reduce((accumulator: number, current: string) => {
        let first = "",
          last = "";
        const chars = current.split("");
        for (let i in chars) {
          const char = chars[i];
          if (numbers.indexOf(char) > -1) {
            if (first === "") {
              first = char;
            } else {
              last = char;
            }
          }
        }
        if (last === "") {
          last = first;
        }
        return accumulator + parseInt(first + last);
      }, 0)
      .toString(10);
  },
  Part2: (input: Array<string>): string => {
    let replace = (
      original: Array<string>,
      i: number,
      map: { [key: string]: string },
      additionalDigit: number
    ): Array<string> => {
      const updated = [...original];
      let search = updated[i - 2] + updated[i - 1] + updated[i];
      if (additionalDigit > 0) search = updated[i - 3] + search;
      if (additionalDigit > 1) search = updated[i - 4] + search;
      const found = map[search];
      if (found !== undefined) updated[i - 2 - additionalDigit] = found;
      return updated;
    };

    const updatedInput: string[] = [];
    for (let i in input) {
      let replaced: string[] = [];
      for (let j in input[i].split("")) {
        replaced.push(input[i][j]);
        if (Number(j) > 1) {
          replaced = replace(
            replaced,
            Number(j),
            {
              one: "1",
              two: "2",
              six: "6",
            },
            0
          );
        }
        if (Number(j) > 2) {
          replaced = replace(
            replaced,
            Number(j),
            {
              four: "4",
              five: "5",
              nine: "9",
            },
            1
          );
        }
        if (Number(j) > 3) {
          replaced = replace(
            replaced,
            Number(j),
            {
              three: "3",
              seven: "7",
              eight: "8",
            },
            2
          );
        }
      }
      updatedInput.push(replaced.join(""));
    }
    return Day1.Part1(updatedInput);
  },
};

Run(Day1);
