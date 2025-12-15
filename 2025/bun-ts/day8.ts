interface JBox {
  x: number;
  y: number;
  z: number;
}

interface Distance {
  startIdx: number;
  endIdx: number;
  distance: number;
}

async function part1() {
  const content = await Bun.file("./test_day8_example.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const rowsSplittedLetters: JBox[] = [];

  rows.forEach((r) => {
    const coords = r.split(",");

    rowsSplittedLetters.push({
      x: parseInt(coords[0]),
      y: parseInt(coords[1]),
      z: parseInt(coords[2]),
    });
  });

  const currentDistanceArr: Distance[] = [];
  for (let i = 0; i < rowsSplittedLetters.length; i++) {
    for (let j = i + 1; j < rowsSplittedLetters.length; j++) {
      const distance = Math.sqrt(
        Math.pow(rowsSplittedLetters[i].x - rowsSplittedLetters[j].x, 2) +
          Math.pow(rowsSplittedLetters[i].y - rowsSplittedLetters[j].y, 2) +
          Math.pow(rowsSplittedLetters[i].z - rowsSplittedLetters[j].z, 2),
      );

      currentDistanceArr.push({ startIdx: i, endIdx: j, distance });
    }
  }
  const test = currentDistanceArr.sort((a, b) => b.distance - a.distance);
  console.log(test);
  const top10 = currentDistanceArr
    .sort((a, b) => a.distance - b.distance)
    .slice(0, 10);
  // .sort((a, b) => a.startIdx - b.startIdx);
  const arrTop10: number[][] = [];
  top10.forEach((a) => arrTop10.push([a.startIdx, a.endIdx]));

  for (let i = 0; i < arrTop10.length; i++) {
    for (let k = i + 1; k < arrTop10.length; k++) {
      const intersection = arrTop10[k].filter((x) => arrTop10[i].includes(x));
      if (intersection!.length > 0) {
        const difference = arrTop10[k]?.filter((x) => !arrTop10[i].includes(x));

        if (difference!.length > 0) {
          arrTop10[i].push(...difference!);
          arrTop10.splice(k, 1);
          i = 0;
        } else {
          arrTop10.splice(k, 1);
          i = 0;
        }
      }
    }
  }

  arrTop10.sort((a, b) => b.length - a.length);
  console.log(arrTop10[0].length * arrTop10[1].length * arrTop10[2].length);
}

// part1();
async function part2() {
  const content = await Bun.file("./test_day8_example.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const rowsSplittedLetters: JBox[] = [];

  rows.forEach((r) => {
    const coords = r.split(",");

    rowsSplittedLetters.push({
      x: parseInt(coords[0]),
      y: parseInt(coords[1]),
      z: parseInt(coords[2]),
    });
  });

  const currentDistanceArr: Distance[] = [];
  for (let i = 0; i < rowsSplittedLetters.length; i++) {
    for (let j = i + 1; j < rowsSplittedLetters.length; j++) {
      const distance = Math.sqrt(
        Math.pow(rowsSplittedLetters[i].x - rowsSplittedLetters[j].x, 2) +
          Math.pow(rowsSplittedLetters[i].y - rowsSplittedLetters[j].y, 2) +
          Math.pow(rowsSplittedLetters[i].z - rowsSplittedLetters[j].z, 2),
      );

      currentDistanceArr.push({ startIdx: i, endIdx: j, distance });
    }
  }
  const top10 = currentDistanceArr
    .sort((a, b) => a.distance - b.distance)
    .slice(0, 10);
  const arrTop10: number[][] = [];

  top10.forEach((a) => arrTop10.push([a.startIdx, a.endIdx]));
  for (let i = 0; i < arrTop10.length; i++) {
    for (let k = i + 1; k < arrTop10.length; k++) {
      const intersection = arrTop10[k].filter((x) => arrTop10[i].includes(x));
      if (intersection!.length > 0) {
        const difference = arrTop10[k]?.filter((x) => !arrTop10[i].includes(x));
        if (difference!.length > 0) {
          arrTop10[i].push(...difference);
          arrTop10.splice(k, 1);
          i = 0;
        } else {
          arrTop10.splice(k, 1);
          i = 0;
        }
      }
    }
  }

  const topRest = currentDistanceArr
    .sort((a, b) => a.distance - b.distance)
    .slice(10);

  console.log(arrTop10);
  arrTop10.sort((a, b) => b.length - a.length);
  console.log(arrTop10[0].length * arrTop10[1].length * arrTop10[2].length);
}

part2();
