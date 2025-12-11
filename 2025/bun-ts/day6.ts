async function part1() {
  const content = await Bun.file("./test_day6.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const cols: string[][] = [];
  rows.forEach((a) => cols.push(a.split(" ").filter((el) => el !== "")));

  // console.log(cols);

  const results: string[][] = [];
  for (let i = 0; i < cols[0].length; i++) {
    const smallArr: string[] = [];
    for (let j = 0; j < cols.length; j++) {
      smallArr.push(cols[j][i]);
    }
    results.push(smallArr);
  }

  let result = 0;
  for (let i = 0; i < results.length; i++) {
    const operator = results[i].pop();
    let localResult = 0;
    for (let j = 0; j < results[i].length; j++) {
      // console.log(results[i][j]);
      if (operator == "*") {
        localResult =
          (localResult == 0 ? 1 : localResult) * parseInt(results[i][j]);
        // console.log(localResult);
      } else {
        localResult += parseInt(results[i][j]);
      }
    }
    result += localResult;
  }
  console.log(result);
}

interface Operator {
  operator: string;
  arr: number[];
}

// part1();
async function part2() {
  const content = await Bun.file("./test_day6.txt").text();
  const rows = content.split("\n").filter((ifilter) => ifilter !== "");
  const cols: string[][] = [];

  rows.forEach((a) => cols.push(a.split("")));

  let currentMax = 0;
  for (let i = 0; i < cols.length; i++) {
    if (currentMax < cols[i].length) {
      currentMax = cols[i].length;
    }
  }
  cols.forEach((a, i) => {
    const padding: string[] = Array(currentMax - a.length).fill(" ");
    cols[i] = cols[i].concat(padding);
  });

  // console.log(currentMax);

  let smallArr: number[] = [];
  const operatorMap: Operator[] = [];
  for (let c = currentMax - 1; c >= 0; c--) {
    let currNum = "";
    for (let z = 0; z < cols.length; z++) {
      if (cols[z][c] !== "") {
        currNum += cols[z][c];
      }
    }
    // console.log(currNum, "new line", currNum[currNum.length - 1]);
    if (currNum[currNum.length - 1] !== " ") {
      const operator = currNum[currNum.length - 1];
      currNum = currNum.slice(0, currNum.length - 1);
      smallArr.push(parseInt(currNum));
      const operatorIntance: Operator = {
        operator,
        arr: smallArr.filter((a) => !isNaN(a)),
      };
      operatorMap.push(operatorIntance);
      smallArr = [];
    } else {
      smallArr.push(parseInt(currNum));
    }
  }

  let result = 0;
  for (let i = 0; i < operatorMap.length; i++) {
    const oper = operatorMap[i].operator;
    const localResult = operatorMap[i].arr.reduce((a, b) => {
      if (oper == "+") {
        a += b;
      } else {
        a = (a == 0 ? 1 : a) * b;
      }
      return a;
    });
    result += localResult;
  }
  console.log(result);
  // const results: string[][] = [];
  // for (let i = 0; i < cols[0].length; i++) {
  //   const smallArr: string[] = [];
  //   for (let j = 0; j < cols.length; j++) {
  //     smallArr.push(cols[j][i]);
  //   }
  //   results.push(smallArr);
  // }
  //
  // console.log(results);
  //
  // let result = 0;
  // for (let i = 0; i < results.length; i++) {
  //   const operator = results[i].pop();
  //   let localResult = 0;
  //   for (let j = 0; j < results[i].length; j++) {
  //     // console.log(results[i][j]);
  //     if (operator == "*") {
  //       localResult =
  //         (localResult == 0 ? 1 : localResult) * parseInt(results[i][j]);
  //       // console.log(localResult);
  //     } else {
  //       localResult += parseInt(results[i][j]);
  //     }
  //   }
  //   result += localResult;
  // }
  // console.log(results);
}

part2();
