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

  const pointArr = points;

  const segments = [];
  for (let i = 0; i < pointArr.length; i++) {
    const start = pointArr[i];
    const end = pointArr[(i + 1) % pointArr.length]; // wrap around to close polygon
    segments.push([start, end]);
  }
  const rectangleCandidates = [];
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
  let rectPoints = [
    [
      { x: validRect.b.x, y: validRect.a.y },
      validRect.a,
      { x: validRect.a.x, y: validRect.b.y },
      validRect.b,
    ],
  ];

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
