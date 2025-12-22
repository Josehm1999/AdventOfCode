interface Machine {
  lightDiagram: string[];
  buttonWiring: number[][];
  joltage: number[];
}

interface MCombinations {
  lightDiagramCombination: string;
  buttonCombination: number[][][];
}

async function part1() {
  const content = await Bun.file("./test_day10_example.txt").text();
  // console.log(content.split('\n'))
  const rows = content.split("\n").filter((x) => x.length > 0);

  const machines: Machine[] = rows.map((x) => {
    const part = x.split(" ");

    const lightDiagram = part[0].replace("[", "").replace("]", "").split("");
    const buttonWiring = part
      .slice(1, part.length - 1)
      .map((x) =>
        x.split(",").map((y) => parseInt(y.replace("(", "").replace(")", ""))),
      );
    const joltage = part[part.length - 1]
      .split(",")
      .map((x) => parseInt(x.replace("{", "").replace("}", "")));

    return {
      lightDiagram,
      buttonWiring,
      joltage,
    };
  });

  const buttonPresses: Map<string, number[][][]> = new Map();
  machines.forEach((machine) => {
    console.log("------");
    machine.buttonWiring.forEach((_, i) => {
      for (let j = 1; j < machine.buttonWiring.length + 1; j++) {
        // console.log(machine.buttonWiring.slice(i, i + j));
        // buttonPresses.push(machine.buttonWiring.slice(i, i + j));
        if (buttonPresses.has(machine.lightDiagram.join(""))) {
          buttonPresses
            .get(machine.lightDiagram.join(""))!
            .push(machine.buttonWiring.slice(i, i + j));
        } else {
          buttonPresses.set(machine.lightDiagram.join(""), [
            machine.buttonWiring.slice(i, i + j),
          ]);
        }
      }
    });
  });
  console.log(buttonPresses);
}

part1();

async function part2() {
  const content = await Bun.file("./test_day10_example.txt").text();
}

// part2();
