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

  // const test = Math.sqrt(
  //   Math.pow(rowsSplittedLetters[0].x - rowsSplittedLetters[1].x, 2) +
  //     Math.pow(rowsSplittedLetters[0].y - rowsSplittedLetters[1].y, 2) +
  //     Math.pow(rowsSplittedLetters[0].z - rowsSplittedLetters[1].z, 2),
  // );

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
  currentDistanceArr.sort((a, b) => a.distance - b.distance);
  // .sort((a, b) => a.startIdx - b.startIdx);

  console.log(currentDistanceArr);
}

part1();
// async function part2() {
//   const content = await Bun.file("./test_day7.txt").text();
//   const rows = content.split("\n").filter((ifilter) => ifilter !== "");
//   const rowsSplittedLetters: string[][] = [];
//
//   rows.forEach((r) => rowsSplittedLetters.push(r.split("")));
//
//   let startingPoint: Point | null = null;
//   for (let j = 0; j < rowsSplittedLetters[0].length; j++) {
//     const currLetter = rowsSplittedLetters[0][j];
//     if (currLetter == "S") {
//       startingPoint = { row: 0, col: j };
//       break;
//     }
//   }
//
//   function recurse(point: Point, memo?: Map<string, number>) {
//     memo = memo || new Map();
//
//     const key1 = `${point.row}-${point.col}`;
//     if (memo.has(key1)) {
//       return memo.get(key1);
//     }
//
//     if (point.row > rowsSplittedLetters.length - 1) {
//       return 1;
//     }
//
//     if (
//       rowsSplittedLetters[point.row][point.col] == "." ||
//       rowsSplittedLetters[point.row][point.col] == "S"
//     ) {
//       return recurse({ row: point.row + 1, col: point.col }, memo);
//     } else if (rowsSplittedLetters[point.row][point.col] == "^") {
//       const val =
//         recurse({ row: point.row, col: point.col - 1 }, memo) +
//         recurse({ row: point.row, col: point.col + 1 }, memo);
//       const key2 = `${point.row}-${point.col}`;
//       memo.set(key2, val);
//       return val;
//     }
//   }
//   // console.log(startingPoint);
//   const counter = recurse(startingPoint!);
//   console.log(counter);
// }
//
// part2();
