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

// part1();
async function part2() {
  const content = await Bun.file("./test_day7_example.txt").text();
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
        }

        if (rowsSplittedLetters[i + 1][j] == "^") {
          rowsSplittedLetters[i + 1][j - 1] = "|";
          rowsSplittedLetters[i + 1][j + 1] = "|";
        }
      }

      if (currLetter == "^") {
        if (rowsSplittedLetters[i - 1][j] == "|") {
          rowsSplittedLetters[i][j - 1] = "|";
          rowsSplittedLetters[i][j + 1] = "|";
        }
      }

      if (currLetter == ".") {
        if (i > 0 && rowsSplittedLetters[i - 1][j] == "|") {
          rowsSplittedLetters[i][j] = "|";
        }
      }
    }
  }

  for (let i = 0; i < rowsSplittedLetters.length; i++) {
    console.log(rowsSplittedLetters[i].join(""));
    for (let j = 0; j < rowsSplittedLetters[i].length; j++) {
      if (
        // i != 2 &&
        rowsSplittedLetters[i][j] == "^" &&
        rowsSplittedLetters[i - 1][j] == "|"
      ) {
        if (rowsSplittedLetters[i][j - 1] == "|") {
          counter++;
        }

        if (rowsSplittedLetters[i][j + 1] == "|") {
          counter++;
        }
      }
    }
  }

  // console.log(rowsSplittedLetters, counter);
}

part2();
