interface Machine {
  lightDiagram: string[];
  buttonWiring: number[][];
  joltage: number[];
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

  machines.forEach((machine) => {
    console.log("------");

    machine.buttonWiring.forEach((_, i) => {
      for (let j = 1; j < machine.buttonWiring.length + 1; j++) {
        console.log(machine.buttonWiring.slice(i, i + j));
        // for (let k = i + 1; k < machine.buttonWiring.length; k++) {
        //   console.log(machine.buttonWiring.slice(i, k + j));
        // }
      }
    });
  });
}

part1();

async function part2() {
  const content = await Bun.file("./test_day10_example.txt").text();
}

// part2();
