async function part1() {
  const content = await Bun.file("./test_day4.txt").text();
  const rows = content.split("\n");

  const arr: string[][] = [];

  for (let k = 0; k < rows.length - 1; k++) {
    arr.push(rows[k].split(""));
  }

  let result = 0;
  for (let i = 0; i < arr.length; i++) {
    for (let j = 0; j < arr[i].length; j++) {
      if (arr[i][j] == "@") {
        // console.log(cols[j]);
        let adyacentCount = 0;

        // j - 1;
        if (j > 0 && arr[i][j - 1] == "@") {
          adyacentCount++;
        }

        // j + 1;
        if (j < arr[i].length && arr[i][j + 1] == "@") {
          adyacentCount++;
        }

        // i - 1;
        if (i > 0 && arr[i - 1][j] == "@") {
          adyacentCount++;
        }

        // i + 1;
        if (i < arr.length - 1 && arr[i + 1][j] == "@") {
          adyacentCount++;
        }

        // (i + 1, j + 1, j - 1);
        if (
          i < arr.length - 1 &&
          j < arr[i].length &&
          arr[i + 1][j + 1] == "@"
        ) {
          adyacentCount++;
        }

        if (i < arr.length - 1 && j > 0 && arr[i + 1][j - 1] == "@") {
          adyacentCount++;
        }
        // (i - 1, j + 1, j - 1);
        if (i > 0 && j < arr[i].length && arr[i - 1][j + 1] == "@") {
          adyacentCount++;
        }

        if (i > 0 && j > 0 && arr[i - 1][j - 1] == "@") {
          adyacentCount++;
        }

        if (adyacentCount < 4) {
          result++;
        }
      }
    }
  }

  console.log(result);
}

// part1();

async function part2() {
  const content = await Bun.file("./test_day4.txt").text();
  const rows = content.split("\n");

  const arr: string[][] = [];

  for (let k = 0; k < rows.length - 1; k++) {
    arr.push(rows[k].split(""));
  }

  let result = 0;
  for (let i = 0; i < arr.length; i++) {
    for (let j = 0; j < arr[i].length; j++) {
      if (arr[i][j] == "@") {
        // console.log(cols[j]);
        let adyacentCount = 0;

        // j - 1;
        if (j > 0 && arr[i][j - 1] == "@") {
          adyacentCount++;
        }

        // j + 1;
        if (j < arr[i].length && arr[i][j + 1] == "@") {
          adyacentCount++;
        }

        // i - 1;
        if (i > 0 && arr[i - 1][j] == "@") {
          adyacentCount++;
        }

        // i + 1;
        if (i < arr.length - 1 && arr[i + 1][j] == "@") {
          adyacentCount++;
        }

        // (i + 1, j + 1, j - 1);
        if (
          i < arr.length - 1 &&
          j < arr[i].length &&
          arr[i + 1][j + 1] == "@"
        ) {
          adyacentCount++;
        }

        if (i < arr.length - 1 && j > 0 && arr[i + 1][j - 1] == "@") {
          adyacentCount++;
        }
        // (i - 1, j + 1, j - 1);
        if (i > 0 && j < arr[i].length && arr[i - 1][j + 1] == "@") {
          adyacentCount++;
        }

        if (i > 0 && j > 0 && arr[i - 1][j - 1] == "@") {
          adyacentCount++;
        }

        if (adyacentCount < 4) {
          result++;
          arr[i][j] = ".";
          i = 0;
          j = 0;
        }
      }
    }
  }

  console.log(result);
}

part2();
