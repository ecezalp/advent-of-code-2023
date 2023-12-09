import { Run, Day } from "../common";

type Location = {
  [key: string]: {
    L: string;
    R: string;
    [key: string]: string;
  };
};

const Day7: Day = {
  Part1: (input: Array<string>): string => {
    let directions: Array<string> = [];
    let locations: Location = {};
    let start: string = "";
    for (let i = 0; i < input.length; i++) {
      if (i == 0) {
        directions = input[i].split("");
      }
      if (i > 1) {
        let name = input[i].slice(0, 3);
        if (i == 2) {
          start = name;
        }
        let left = input[i].slice(7, 10);
        let right = input[i].slice(12, 15);
        locations[name] = {
          L: left,
          R: right,
        };
      }
    }

    let current = "AAA";
    let times = 0;
    for (let i = 0; i < directions.length; i++) {
      if (current != "ZZZ") {
        current = locations[current][directions[i]];
      } else {
        return ((times * directions.length) + i).toString();
        break;
      }
      if (i == directions.length - 1) {
        times++;
        i = -1;
      }
    }

    return "";
  },
  Part2: (input: Array<string>): string => {
    let directions: Array<string> = [];
    let locations: Location = {};
    let start: string = "";
    for (let i = 0; i < input.length; i++) {
      if (i == 0) {
        directions = input[i].split("");
      }
      if (i > 1) {
        let name = input[i].slice(0, 3);
        if (i == 2) {
          start = name;
        }
        let left = input[i].slice(7, 10);
        let right = input[i].slice(12, 15);
        locations[name] = {
          L: left,
          R: right,
        };
      }
    }

    let currents = [];
    for (let i = 0; i < Object.keys(locations).length; i++) {
      if (Object.keys(locations)[i][2]== "A") {
        currents.push(Object.keys(locations)[i])
      }
    }
    console.log(currents)
    console.log(currents.length)
    console.log(currents.length)

    let items = []
    for (let i in currents) {
      let times = 0;
      let current = currents[i]
      console.log(current)
      for (let i = 0; i < directions.length; i++) {
        if (current[2] != "Z") {
          current = locations[current][directions[i]];
        } else {
          let item = (times * directions.length) + i
          items.push(item)
          break;
        }
        if (i == directions.length - 1) {
          times++;
          i = -1;
        }
      }
    }
    let lcm = items[0]
    for (let i in items) {
      if (i == "0") {
        lcm = items[i]
      } else {
        lcm = findLCM(lcm, items[i])
      }
    }

    return lcm.toString();
  },
};

function findGCD(a: number, b: number): number {
  return b === 0 ? a : findGCD(b, a % b);
}

// Function to find the least common multiple (LCM) of two numbers
function findLCM(a: number, b: number): number {
  return (a * b) / findGCD(a, b);
}

Run(Day7);
