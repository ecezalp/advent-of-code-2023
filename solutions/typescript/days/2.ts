import { Run, Day } from "../common";

const Day2: Day = {
  Part1: (input: Array<string>): string => {
    return input
      .reduce((accumulator: number, line: string) => {
        const game = line.split(": ");
        const id = parseInt(game[0].slice(4));
        const rounds = game[1].split("; ");
        for (let i = 0; i < rounds.length; i++) {
          const colorPairs = rounds[i].split(", ");
          for (let j = 0; j < colorPairs.length; j++) {
            const colorTuple = colorPairs[j].split(" ");
            if (
              (colorTuple[1] == "red" && parseInt(colorTuple[0]) > 12) ||
              (colorTuple[1] == "blue" && parseInt(colorTuple[0]) > 14) ||
              (colorTuple[1] == "green" && parseInt(colorTuple[0]) > 13)
            ) {
              return accumulator;
            }
          }
        }
        return accumulator + id;
      }, 0)
      .toString(10);
  },
  Part2: (input: Array<string>): string => {
    return input
      .reduce((accumulator: number, line: string) => {
        const game = line.split(": ");
        const rounds = game[1].split("; ");
        let redMax = 0,
          blueMax = 0,
          greenMax = 0;
        for (let i = 0; i < rounds.length; i++) {
          const colorPairs = rounds[i].split(", ");
          for (let j = 0; j < colorPairs.length; j++) {
            const colorTuple = colorPairs[j].split(" ");
            const value = parseInt(colorTuple[0]);
            if (colorTuple[1] == "red" && value > redMax) {
              redMax = value;
            }
            if (colorTuple[1] == "blue" && value > blueMax) {
              blueMax = value;
            }
            if (colorTuple[1] == "green" && value > greenMax) {
              greenMax = value;
            }
          }
        }
        return accumulator + redMax * greenMax * blueMax;
      }, 0)
      .toString(10);
  },
};

Run(Day2);
