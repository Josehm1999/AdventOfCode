async function part1() {
  const content = await Bun.file("./test_day7.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const rowsSplittedLetters: string[][] = [];

  rows.forEach((r) => rowsSplittedLetters.push(r.split("")));

  let counter = 0;
  for (let i = 0; i < rowsSplittedLetters.length; i++) {
    for (let j = 0; j < rowsSplittedLetters[i].length; j++) {
      // console.log(rowsSplittedLetters[i][j]);
      const currLetter = rowsSplittedLetters[i][j];
      if (currLetter == "S") {
        if (rowsSplittedLetters[i + 1][j] == ".") {
          rowsSplittedLetters[i + 1][j] = "|";
          // counter++;
        }

        if (rowsSplittedLetters[i + 1][j] == "^") {
          rowsSplittedLetters[i + 1][j - 1] = "|";
          rowsSplittedLetters[i + 1][j + 1] = "|";
          // counter += 2;
        }
      }

      if (currLetter == "^") {
        if (rowsSplittedLetters[i - 1][j] == "|") {
          rowsSplittedLetters[i][j - 1] = "|";
          rowsSplittedLetters[i][j + 1] = "|";
          counter += 1;
        }
      }

      if (currLetter == ".") {
        if (i > 0 && rowsSplittedLetters[i - 1][j] == "|") {
          rowsSplittedLetters[i][j] = "|";
          // counter++;
        }
      }
    }
  }

  console.log(rowsSplittedLetters, counter);
}

interface Point {
  row: number;
  col: number;
}

// part1();
async function part2() {
  const content = await Bun.file("./test_day7.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const rowsSplittedLetters: string[][] = [];

  rows.forEach((r) => rowsSplittedLetters.push(r.split("")));

  let startingPoint: Point | null = null;
  for (let j = 0; j < rowsSplittedLetters[0].length; j++) {
    const currLetter = rowsSplittedLetters[0][j];
    if (currLetter == "S") {
      startingPoint = { row: 0, col: j };
      break;
    }
  }

  function recurse(point: Point, memo?: Map<string, number>) {
    memo = memo || new Map();

    const key1 = `${point.row}-${point.col}`;
    if (memo.has(key1)) {
      return memo.get(key1);
    }

    if (point.row > rowsSplittedLetters.length - 1) {
      return 1;
    }

    if (
      rowsSplittedLetters[point.row][point.col] == "." ||
      rowsSplittedLetters[point.row][point.col] == "S"
    ) {
      return recurse({ row: point.row + 1, col: point.col }, memo);
    } else if (rowsSplittedLetters[point.row][point.col] == "^") {
      const val =
        recurse({ row: point.row, col: point.col - 1 }, memo) +
        recurse({ row: point.row, col: point.col + 1 }, memo);
      const key2 = `${point.row}-${point.col}`;
      memo.set(key2, val);
      return val;
    }
  }
  // console.log(startingPoint);
  const counter = recurse(startingPoint!);
  console.log(counter);
}

part2();
