async function part1() {
  const content = await Bun.file("./test_day2.txt").text();

  const ranges = content.split(",");

  const invalidIds: number[] = [];
  for (let i = 0; i < ranges.length; i++) {
    const parts = ranges[i].split("-");
    const left = parseInt(parts[0]);
    const right = parseInt(parts[1]);

    for (let start = left; start <= right; start++) {
      if (start.toString().length % 2 == 0) {
        const midPoint = start.toString().length / 2;
        const leftSection = start
          .toString()
          .split("")
          .slice(0, midPoint)
          .join("");
        const rightSection = start
          .toString()
          .split("")
          .slice(midPoint, start.toString().split("").length)
          .join("");

        if (leftSection == rightSection) {
          invalidIds.push(start);
        }
      }
    }
  }

  console.log(invalidIds.reduce((a, b) => a + b));
}

async function part2() {
  const content = await Bun.file("./test_day2.txt").text();

  const ranges = content.split(",");

  const invalidIds: number[] = [];
  for (let i = 0; i < ranges.length; i++) {
    const parts = ranges[i].split("-");
    const left = parseInt(parts[0]);
    const right = parseInt(parts[1]);

    for (let start = left; start <= right; start++) {
      // console.log(start, right);
      for (let divider = 1; divider <= 9; divider++) {
        const section: string[] = [];
        if (start.toString().length % divider == 0) {
          for (
            let start1 = 0;
            start1 < start.toString().length;
            start1 = start1 + divider
          ) {
            section.push(
              start
                .toString()
                .split("")
                .slice(start1, start1 + divider)
                .join("")
            );
          }
          if (section.length > 1) {
            if ([...new Set(section)].length === 1) {
              invalidIds.push(start);
              break;
            }
          }
        }
      }
    }
  }

  console.log(invalidIds.reduce((a, b) => a + b));
}

part1();
part2();

export {};
