interface MPoint {
  x: number;
  y: number;
}

async function part1() {
  const content = await Bun.file("./test_day9.txt").text();
  const pointArr: MPoint[] = [];
  content
    .split("\n")
    .filter((ifilter) => ifilter !== "")
    .forEach((x) => {
      const splits = x.split(",");
      pointArr.push({ x: parseInt(splits[1]), y: parseInt(splits[0]) });
    });

  const areaByPairs: Map<string, number> = new Map();
  for (let i = 0; i < pointArr.length; i++) {
    for (let j = i + 1; j < pointArr.length; j++) {
      const key = `${i}-${j}`;
      if (pointArr[i].x == pointArr[j].x) {
        areaByPairs.set(
          key,
          (pointArr[i].y > pointArr[j].y
            ? pointArr[i].y - pointArr[j].y
            : pointArr[j].y - pointArr[i].y) + 2,
        );
        continue;
      }

      if (pointArr[i].y == pointArr[j].y) {
        areaByPairs.set(
          key,
          (pointArr[i].x > pointArr[j].x
            ? pointArr[i].x - pointArr[j].x
            : pointArr[j].x - pointArr[i].x) + 2,
        );

        continue;
      }

      if (
        (pointArr[i].x > pointArr[j].x
          ? pointArr[i].x - pointArr[j].x
          : pointArr[j].x - pointArr[i].x) ==
        (pointArr[i].y > pointArr[j].y
          ? pointArr[i].y - pointArr[j].y
          : pointArr[j].y - pointArr[i].y)
      ) {
        continue;
      } else {
        // console.log(
        //   key,
        //   pointArr[i],
        //   pointArr[j],
        //   ((pointArr[i].x > pointArr[j].x
        //     ? pointArr[i].x - pointArr[j].x
        //     : pointArr[j].x - pointArr[i].x) +
        //     1) *
        //     ((pointArr[i].y > pointArr[j].y
        //       ? pointArr[i].y - pointArr[j].y
        //       : pointArr[j].y - pointArr[i].y) +
        //       1),
        // );
        areaByPairs.set(
          key,
          ((pointArr[i].x > pointArr[j].x
            ? pointArr[i].x - pointArr[j].x
            : pointArr[j].x - pointArr[i].x) +
            1) *
            ((pointArr[i].y > pointArr[j].y
              ? pointArr[i].y - pointArr[j].y
              : pointArr[j].y - pointArr[i].y) +
              1),
        );
      }
    }
  }
  let area = 0;
  for (let val of areaByPairs.values()) {
    if (area < val) {
      area = val;
    }
  }
  console.log(area);
}

// part1();
async function part2() {
  const content = await Bun.file("./test_day9.txt").text();
  const pointArr: MPoint[] = [];
  content
    .split("\n")
    .filter((ifilter) => ifilter !== "")
    .forEach((x) => {
      const splits = x.split(",");
      pointArr.push({ x: parseInt(splits[1]), y: parseInt(splits[0]) });
    });

  const areaByPairs: Map<string, number> = new Map();
  for (let i = 0; i < pointArr.length; i++) {
    for (let j = i + 1; j < pointArr.length; j++) {
      const key = `${i}-${j}`;
      if (pointArr[i].x == pointArr[j].x) {
        areaByPairs.set(
          key,
          (pointArr[i].y > pointArr[j].y
            ? pointArr[i].y - pointArr[j].y
            : pointArr[j].y - pointArr[i].y) + 2,
        );
        continue;
      }

      if (pointArr[i].y == pointArr[j].y) {
        areaByPairs.set(
          key,
          (pointArr[i].x > pointArr[j].x
            ? pointArr[i].x - pointArr[j].x
            : pointArr[j].x - pointArr[i].x) + 2,
        );

        continue;
      }

      if (
        (pointArr[i].x > pointArr[j].x
          ? pointArr[i].x - pointArr[j].x
          : pointArr[j].x - pointArr[i].x) ==
        (pointArr[i].y > pointArr[j].y
          ? pointArr[i].y - pointArr[j].y
          : pointArr[j].y - pointArr[i].y)
      ) {
        continue;
      } else {
        // console.log(
        //   key,
        //   pointArr[i],
        //   pointArr[j],
        //   ((pointArr[i].x > pointArr[j].x
        //     ? pointArr[i].x - pointArr[j].x
        //     : pointArr[j].x - pointArr[i].x) +
        //     1) *
        //     ((pointArr[i].y > pointArr[j].y
        //       ? pointArr[i].y - pointArr[j].y
        //       : pointArr[j].y - pointArr[i].y) +
        //       1),
        // );
        //

        const newPointA: MPoint = { x: pointArr[i].x, y: pointArr[j].y };
        const newPointB: MPoint = { x: pointArr[j].x, y: pointArr[i].y };

        // console.log("A", newPointA, "B", newPointB);
        let pointAflag = false;
        let pointBflag = false;
        // 4 variaciones para cada ezquina
        if (pointArr[i].x > pointArr[j].x && pointArr[i].y > pointArr[j].y) {
          // console.log(1)
          if (pointArr[i].x == 5 && pointArr[i].y == 9) {
            console.log(pointArr[i], pointArr[j]);
          }
          //i esta a abajo y a la derecha
          // newPointA esta abajo y a la izquierda
          pointAflag = pointArr.some(
            (p) =>
              (p.x == newPointA.x && p.y == newPointA.y) ||
              (p.x > newPointA.x && p.y < newPointA.y),
          );

          // newPointB esta arriba y a la derecha
          pointBflag = pointArr.some(
            (p) =>
              (p.x == newPointB.x && p.y == newPointB.y) ||
              (p.x < newPointB.x && p.y > newPointB.y),
          );
        }

        if (pointArr[j].x > pointArr[i].x && pointArr[j].y > pointArr[i].y) {
          if (pointArr[i].x == 5 && pointArr[i].y == 9) {
            console.log(pointArr[i], pointArr[j]);
          }
          //i esta arriba y a la izquierda
          // newPointA esta arriba y a la derecha
          pointAflag = pointArr.some(
            (p) => p.x <= newPointA.x && p.y >= newPointA.y,
            // (p.x < newPointA.x && p.y > newPointA.y),
          );
          // newPointB esta abajo y a la izquierda
          pointBflag = pointArr.some(
            (p) => p.x >= newPointB.x && p.y <= newPointB.y,
            // (p.x > newPointB.x && p.y < newPointB.y),
          );
        }

        if (pointArr[j].x > pointArr[i].x && pointArr[i].y > pointArr[j].y) {
          if (pointArr[i].x == 5 && pointArr[i].y == 9) {
            console.log(pointArr[i], pointArr[j]);
          }
          //i esta arriba y a la derecha
          // newPointA esta arriba y a la izquierda
          pointAflag = pointArr.some(
            (p) => p.x <= newPointA.x && p.y <= newPointA.y,
            // (p.x < newPointA.x && p.y < newPointA.y),
          );
          // newPointB esta abajo y a la derecha
          pointBflag = pointArr.some(
            (p) => p.x >= newPointB.x && p.y >= newPointB.y,
            // (p.x > newPointB.x && p.y > newPointB.y),
          );
        }

        if (pointArr[i].x > pointArr[j].x && pointArr[j].y > pointArr[i].y) {
          if (pointArr[i].x == 5 && pointArr[i].y == 9) {
            console.log(pointArr[i], pointArr[j]);
          }

          // i esta abajo y a la izquierda
          // newPointA esta abajo y a la derecha
          pointAflag = pointArr.some(
            (p) => p.x >= newPointA.x && p.y >= newPointA.y,
            // (p.x > newPointA.x && p.y > newPointA.y),
          );
          // newPointB esta arriba y a la izquierda
          pointBflag = pointArr.some(
            (p) => p.x <= newPointB.x && p.y <= newPointB.y,
            // (p.x < newPointB.x && p.y < newPointB.y),
          );
        }

        // if (pointAflag && pointBflag) {
        //   console.log(pointArr[i], pointArr[j]);
        // }
        if (pointAflag && pointBflag) {
          areaByPairs.set(
            key,
            ((pointArr[i].x > pointArr[j].x
              ? pointArr[i].x - pointArr[j].x
              : pointArr[j].x - pointArr[i].x) +
              1) *
              ((pointArr[i].y > pointArr[j].y
                ? pointArr[i].y - pointArr[j].y
                : pointArr[j].y - pointArr[i].y) +
                1),
          );
        }
      }
    }
  }
  let area = 0;
  for (let [key, val] of areaByPairs.entries()) {
    // console.log(key, val);
    if (area < val) {
      area = val;
    }
  }
  console.log(area);
}

part2();
