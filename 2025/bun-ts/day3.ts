async function part1() {
  const content = await Bun.file("./test_day3.txt").text();

  const ranges = content.split("\n");

  let result = 0;
  for (let i = 0; i < ranges.length - 1; i++) {
    const line = ranges[i].split("");
    const highestNumber = line
      .slice(0, line.length - 1)
      .sort((a, b) => parseInt(a) - parseInt(b))
      .reverse()[0];
    const firstNumberIdx = ranges[i].indexOf(highestNumber);
    const secondNumberHighest = line
      .slice(firstNumberIdx + 1, line.length)
      .sort((a, b) => parseInt(a) - parseInt(b))
      .reverse()[0];

    console.log(line[firstNumberIdx] + secondNumberHighest);
    result += parseInt(line[firstNumberIdx] + secondNumberHighest);
  }

  console.log(result);
}

// part1();
async function part2() {
  const content = await Bun.file("./test_day3.txt").text();

  const ranges = content.split("\n");

  let result = 0;
  for (let i = 0; i < ranges.length - 1; i++) {
    const line = ranges[i].split("");
    // console.log(line);
    let currConcatenatedVal: string[] = [];
    const highestNumber = line
      .slice(0, line.length - 11)
      .sort((a, b) => parseInt(a) - parseInt(b))
      .reverse()[0];

    currConcatenatedVal.push(highestNumber);
    const firstNumberIdx = line.indexOf(highestNumber);

    let currNumberHighest = firstNumberIdx;
    for (let j = 10; j >= 0; j--) {
      const numberHighest = line
        .slice(currNumberHighest + 1, line.length - j)
        .sort((a, b) => parseInt(b) - parseInt(a))[0];

      currNumberHighest =
        currNumberHighest +
        1 +
        line
          .slice(currNumberHighest + 1, line.length - j)
          .indexOf(numberHighest);

      currConcatenatedVal.push(line[currNumberHighest]);
    }

    // console.log(currConcatenatedVal.join(""));
    result += parseInt(currConcatenatedVal.join(""));
  }

  console.log(result);
}
part2();
