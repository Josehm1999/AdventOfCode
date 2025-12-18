const canvas = document.getElementById("myCanvas");
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;
const ctx = canvas.getContext("2d");

const content = await fetch("/test_day9.txt");
let points = [];

function is_inside(arr, newPoint, memo) {
  memo = memo || new Map();

  const key1 = `${newPoint.x}-${newPoint.y}`;
  if (memo.has(key1)) {
    return memo.get(key1);
  }

  let inside = false;
  const x = newPoint.x;
  const y = newPoint.y;
  let p1 = arr[0];
  let p2;
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

function doSegmentsIntersect(A, B, C, D) {
  const o1 = Morientation(A, B, C);
  const o2 = Morientation(A, B, D);
  const o3 = Morientation(C, D, A);
  const o4 = Morientation(C, D, B);
  if (o1 != o2 && o3 != o4) {
    return true;
  }

  if (o1 == 0 && onSegment(A, C, B)) return true;
  if (o2 == 0 && onSegment(A, D, B)) return true;
  if (o3 == 0 && onSegment(C, A, D)) return true;
  if (o4 == 0 && onSegment(C, B, D)) return true;

  return false;
}

function onSegment(p, q, r) {
  return (
    q.x <= Math.max(p.x, r.x) &&
    q.x >= Math.min(p.x, r.x) &&
    q.y <= Math.max(p.y, r.y) &&
    q.y >= Math.min(p.y, r.y)
  );
}

function Morientation(p, q, r) {
  const val = (q.y - p.y) * (r.x - q.x) - (q.x - p.x) * (r.y - q.y);
  if (val == 0) {
    return 0;
  }
  if (val > 0) {
    return 1;
  } else {
    return 2;
  }
}

content.text().then((x) => {
  x.split("\n")
    .filter((ifilter) => ifilter !== "")
    .forEach((x) => {
      const splits = x.split(",");
      points.push({
        x: parseInt(splits[0]),
        y: parseInt(splits[1]),
      });
    });

  let rectPoints = [];
  let currentMax = [];
  let pointArr = points;
  for (let i = 0; i < pointArr.length; i++) {
    for (let j = i + 1; j < pointArr.length; j++) {
      if (pointArr[i].x == pointArr[j].x) {
        continue;
      }

      if (pointArr[i].y == pointArr[j].y) {
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
        const newPointA = { x: pointArr[i].x, y: pointArr[j].y };
        const newPointB = { x: pointArr[j].x, y: pointArr[i].y };

        const isInsideA = is_inside(pointArr, newPointA);
        const isInsideB = is_inside(pointArr, newPointB);

        if (isInsideA && isInsideB) {
          let checkA;
          let checkB;
          let checkC;
          let checkD;
          for (let k = 0; k < pointArr.length - 1; k++) {
            const firstS = pointArr[k];
            const secondS = pointArr[k + 1];

            checkA = doSegmentsIntersect(
              pointArr[i],
              newPointA,
              firstS,
              secondS,
            );
            checkB = doSegmentsIntersect(
              newPointA,
              pointArr[j],
              firstS,
              secondS,
            );
            checkC = doSegmentsIntersect(
              pointArr[j],
              newPointB,
              firstS,
              secondS,
            );
            checkD = doSegmentsIntersect(
              newPointB,
              pointArr[i],
              firstS,
              secondS,
            );
          }

          if (checkA && checkB && checkC && checkD) {
            // console.log("entra", isInsideA.counter);
            const test =
              ((pointArr[i].x > pointArr[j].x
                ? pointArr[i].x - pointArr[j].x
                : pointArr[j].x - pointArr[i].x) +
                1) *
              ((pointArr[i].y > pointArr[j].y
                ? pointArr[i].y - pointArr[j].y
                : pointArr[j].y - pointArr[i].y) +
                1);
            if (currentMax < test) {
              currentMax = test;
              const curr = [pointArr[i], newPointA, pointArr[j], newPointB];
              rectPoints = [];
              rectPoints.push(curr);
            }
            // }
          }
        }
      }
    }
  }
  console.log(points);

  const xs = points.map((p) => p.x);
  const ys = points.map((p) => p.y);
  const minX = Math.min(...xs);
  const maxX = Math.max(...xs);
  const minY = Math.min(...ys);
  const maxY = Math.max(...ys);

  const width = canvas.width;
  const height = canvas.height;
  const padding = 50; // pixels of padding around edges

  // Calculate scale to fit
  const scaleX = (width - 2 * padding) / (maxX - minX);
  const scaleY = (height - 2 * padding) / (maxY - minY);
  const scale = Math.min(scaleX, scaleY); // use smaller scale to fit both dimensions

  // Transform function
  const transform = (p) => ({
    x: (p.x - minX) * scale + padding,
    y: (p.y - minY) * scale + padding,
  });

  // Use transformed points when drawing
  const transformedPoints = points.map(transform);
  // Start drawing the path
  ctx.beginPath();

  // Move to the first point
  ctx.moveTo(transformedPoints[0].x, transformedPoints[0].y);

  // Draw lines to the subsequent points
  for (let i = 1; i < transformedPoints.length; i++) {
    ctx.lineTo(transformedPoints[i].x, transformedPoints[i].y);
  }

  // Close the path (connects the last point to the first point)
  ctx.closePath();

  // Style and fill the polygon
  ctx.fillStyle = "blue";
  ctx.fill();
  ctx.strokeStyle = "black";
  ctx.lineWidth = 2;
  ctx.stroke();

  for (let k = 0; k < rectPoints.length; k++) {
    const curr = rectPoints[k];

    const transformedRect = curr.map(transform);

    // Draw the rectangle
    ctx.beginPath();
    ctx.moveTo(transformedRect[0].x, transformedRect[0].y);
    for (let i = 1; i < transformedRect.length; i++) {
      ctx.lineTo(transformedRect[i].x, transformedRect[i].y);
    }
    ctx.closePath();

    // Style it differently so it stands out
    ctx.fillStyle = "rgba(255, 0, 0, 0.1)"; // semi-transparent red
    ctx.fill();
    ctx.strokeStyle = "red";
    ctx.lineWidth = 3;
    ctx.stroke();
  }
});

// console.log(points);
