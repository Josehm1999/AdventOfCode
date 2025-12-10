async function part1() {
  const content = await Bun.file("./test_day5.txt").text();
  const rows = content.split("\n");
  // console.log(rows);

  let fresh_ranges: string[] = [];
  let available_items: string[] = [];
  let jump_to_items = false;
  for (let j = 0; j < rows.length - 1; j++) {
    if (!jump_to_items && rows[j] != "") {
      fresh_ranges.push(rows[j]);
    }

    if (rows[j] == "") {
      jump_to_items = true;
    }

    if (jump_to_items && rows[j] != "") {
      available_items.push(rows[j]);
    }
  }

  const fresh_items = new Set();
  // console.log(fresh_ranges);
  // console.log(available_items);

  for (let i = 0; i < available_items.length; i++) {
    const item = parseInt(available_items[i]);
    for (let k = 0; k < fresh_ranges.length; k++) {
      const ranges = fresh_ranges[k].split("-");
      const leftRange = parseInt(ranges[0]);
      const rightRange = parseInt(ranges[1]);

      if (item >= leftRange && rightRange >= item) {
        fresh_items.add(item);
      }
    }
  }

  console.log(fresh_items.size);
}

// part1();
//
interface MyRange {
  start: number;
  end: number;
}

async function part2() {
  const content = await Bun.file("./test_day5_example.txt").text();
  const rows = content.split("\n");

  let fresh_ranges: MyRange[] = [];
  let jump_to_items = false;
  for (let j = 0; j < rows.length - 1; j++) {
    if (!jump_to_items && rows[j] != "") {
      const ranges = rows[j].split("-");
      const range: MyRange = {
        start: parseInt(ranges[0]),
        end: parseInt(ranges[1]),
      };
      fresh_ranges.push(range);
    }

    if (rows[j] == "") {
      jump_to_items = true;
    }
  }

  fresh_ranges.sort((a, b) => a.start - b.start);
  for (let k = 0, len = fresh_ranges.length; k < len; k++) {
    if (k >= fresh_ranges.length - 1) {
      break;
    }

    // console.log(fresh_ranges[k], fresh_ranges[k + 1]);
    if (fresh_ranges[k].end >= fresh_ranges[k + 1].start) {
      //overlap
      fresh_ranges[k].end =
        fresh_ranges[k + 1].end >= fresh_ranges[k].end
          ? fresh_ranges[k + 1].end
          : fresh_ranges[k].end;
      fresh_ranges.splice(k + 1, 1);
      fresh_ranges.sort((a, b) => a.start - b.start);
      k = 0;
    }
    // console.log(fresh_ranges);
  }
  console.log(fresh_ranges);
  const result = fresh_ranges.reduce((a, b) => a + (b.end - b.start + 1), 0);
  console.log(result);
}

part2();
