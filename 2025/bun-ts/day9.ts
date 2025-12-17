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
//
function is_inside(arr: MPoint[], newPoint: MPoint) {
  let count = 0;

  for (let i = 0; i < arr.length - 1; i++) {
    const first = arr[i];
    const second = arr[i + 1];

    if (
      ((newPoint.x > first.x && newPoint.x < second.x) ||
        (newPoint.x < first.x && newPoint.x > second.x)) &&
      newPoint.y <
        first.y +
          ((newPoint.x - first.x) / (second.x - first.x)) * (second.y - first.y)
    ) {
      count += 1;
    }
  }

  // si es par significa que salio sino ta dentro
  return count % 2 == 1;
}

async function part2() {
  const content = await Bun.file("./test_day9_example.txt").text();
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

        const isInsideA = is_inside(pointArr, newPointA);
        const isInsideB = is_inside(pointArr, newPointB);
        const isInsideC = is_inside(pointArr, pointArr[i]);
        const isInsideD = is_inside(pointArr, pointArr[j]);

        if (isInsideA && isInsideB && isInsideC && isInsideD) {
          console.log(newPointA, newPointB, pointArr[i], pointArr[j]);
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
