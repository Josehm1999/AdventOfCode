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
function is_inside(
  arr: MPoint[],
  newPoint: MPoint,
  memo?: Map<string, boolean>,
) {
  memo = memo || new Map();

  const key1 = `${newPoint.x}-${newPoint.y}`;
  if (memo.has(key1)) {
    return memo.get(key1);
  }

  let inside = false;
  const x = newPoint.x;
  const y = newPoint.y;
  let p1 = arr[0];
  let p2: MPoint;
  for (let i = 1; i <= arr.length; i++) {
    if (i < arr.length && arr[i].x == x && arr[i].y == y) {
      return true;
    }

    p2 = arr[i % arr.length];

    if (y > Math.min(p1.y, p2.y)) {
      if (y <= Math.max(p1.y, p2.y)) {
        if (x <= Math.max(p1.x, p2.x)) {
          const x_intersection =
            ((y - p1.y) * (p2.x - p1.x)) / (p2.y - p1.y) + p1.x;
          if (p1.x === p2.x || x <= x_intersection) {
            inside = !inside;
          }
        }
      }
    }
    p1 = p2;
  }
  // si es par significa que salio sino ta dentro
  memo.set(key1, inside);
  return inside;
}
async function part2() {
  const content = await Bun.file("./test_day9.txt").text();
  const pointArr: MPoint[] = [];
  content
    .split("\n")
    .filter((ifilter) => ifilter !== "")
    .forEach((x) => {
      const splits = x.split(",");
      pointArr.push({ x: parseInt(splits[0]), y: parseInt(splits[1]) });
    });

  const segments: MPoint[][] = [] as MPoint[][];
  for (let i = 0; i < pointArr.length; i++) {
    const start = pointArr[i];
    const end = pointArr[(i + 1) % pointArr.length]; // wrap around to close polygon
    segments.push([start, end]);
  }
  const rectangleCandidates: { a: MPoint; b: MPoint; area: number }[] = [] as {
    a: MPoint;
    b: MPoint;
    area: number;
  }[];
  for (let i = 0; i < pointArr.length; i++) {
    for (let j = i + 1; j < pointArr.length; j++) {
      const a = pointArr[i];
      const b = pointArr[j];
      const area = (Math.abs(a.x - b.x) + 1) * (Math.abs(a.y - b.y) + 1);
      rectangleCandidates.push({ a, b, area });
    }
  }

  rectangleCandidates.sort((x, y) => y.area - x.area);

  const validRect = rectangleCandidates.find(({ a, b }) => {
    return segments.every((val) => {
      const lineStart = val[0];
      const lineEnd = val[1];
      const leftOfRect = Math.max(a.x, b.x) <= Math.min(lineStart.x, lineEnd.x);
      const rightOfRect =
        Math.min(a.x, b.x) >= Math.max(lineStart.x, lineEnd.x);
      const above = Math.max(a.y, b.y) <= Math.min(lineStart.y, lineEnd.y);
      const below = Math.min(a.y, b.y) >= Math.max(lineStart.y, lineEnd.y);

      return leftOfRect || rightOfRect || above || below;
    });
  });

  const maxArea = validRect ? validRect.area : 0;
  console.log(maxArea);
}

part2();
