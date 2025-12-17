const canvas = document.getElementById("myCanvas");
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;
const ctx = canvas.getContext("2d");

const content = await fetch("/test_day9.txt");
let points = [];

function is_inside(arr, newPoint) {
  let count = 0;

  for (let i = 0; i < arr.length - 1; i++) {
    const first = arr[i];
    const second = arr[i + 1];

    if (
      ((newPoint.y > first.y && newPoint.y < second.y) ||
        (newPoint.y < first.y && newPoint.y > second.y)) &&
      newPoint.x <
        first.x +
          ((newPoint.y - first.y) / (second.y - first.y)) * (second.x - first.x)
    ) {
      count += 1;
    }
  }

  // si es par significa que salio sino ta dentro
  return count % 2 == 1;
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

        // const isInsideA = is_inside(pointArr, newPointA);
        // const isInsideB = is_inside(pointArr, newPointB);
        // const isInsideC = is_inside(pointArr, pointArr[i]);
        // const isInsideD = is_inside(pointArr, pointArr[j]);

        // if (isInsideA && isInsideB && isInsideC && isInsideD) {
        // console.log("entra");
        //single rectangle

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

  // const rectPoints = [
  //   { x: 2, y: 5 }, // top-left
  //   { x: 2, y: 3 }, // bottom-left
  //   { x: 9, y: 3 }, // top-right
  //   { x: 9, y: 5 }, // bottom-right
  // ];
  //
  //
  //
  let largestIdx = 0;
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
