class Program
{

    struct Point
    {
        public int row { set; get; }
        public int col { set; get; }
    }

    static void Main()
    {

        var inputLines = File.ReadAllLines("input.txt");
        var startPoints = new List<Point>();
        var topoMap = new List<List<char>>();
        var path = new List<List<Point>>();
        var seen = new List<List<bool>>();

        for (int i = 0; i < inputLines.Length; i++)
        {
            for (int j = 0; j < inputLines[i].Length; j++)
            {
                if (inputLines[i][j] == '0')
                {
                    // Console.WriteLine($"{i},{j}");
                    var newP = new Point { col = i, row = j };
                    startPoints.Add(newP);
                }
            }

            List<char> charList = new List<char>(inputLines[i].ToCharArray());
            topoMap.Add(charList);
            var tmpSeen = new List<bool>(inputLines[i].Length);
            var tmpPath = new List<Point>(inputLines[i].Length);

            seen.Add(tmpSeen);
            path.Add(tmpPath);
        }

        Dictionary<Point, bool> dups = new Dictionary<Point, bool>();

        var Total = 0;
        foreach (var start in startPoints)
        {
            var results = walk(topoMap, start);
            results.ForEach(i =>
            {
                if (!dups.ContainsKey(i))
                {
                    Console.WriteLine($"Trailhead: ({i.col},{i.row})");
                    dups.Add(i, true);
                }
            });

            Total += dups.Count();
        }
        Console.WriteLine(Total);
    }

    static bool IsOffMap(Point current, int width, int height)
    {
        return (current.row < 0 || current.row >= width || current.col < 0 || current.col >= height);
    }

    static int TopoMapGet(List<List<char>> topoMap, Point point)
    {

        var value = topoMap[point.col][point.row];
        try
        {
            return int.Parse(value.ToString());
        }
        catch (System.Exception)
        {
            return 0;
        }
    }

    static List<Point> walk(List<List<char>> topoMap, Point current)
    {

        var directions = new[] { (-1, 0), (1, 0), (0, -1), (0, 1) };
        Console.WriteLine($"({current.col}, {current.row})");
        // End of trail - good
        if (TopoMapGet(topoMap, current) == 9)
        {
            var result = new List<Point>();
            result.Add(current);
            return result;
        }
        else
        {
            // Recurse
            var result = new List<Point>();
            foreach ((int row_d, int col_d) direction in directions)
            {
                var newPos = new Point { row = current.row + direction.row_d, col = current.col + direction.col_d };
                var isIt = IsOffMap(newPos, topoMap[0].Count(), topoMap.Count());

                // Console.WriteLine($"({current.col}, {current.row}, {TopoMapGet(topoMap, current)}), ({newPos.col}, {newPos.row}, {(!isIt ? TopoMapGet(topoMap, newPos) : false)}), {isIt}");
                if (!isIt && TopoMapGet(topoMap, current) + 1 == TopoMapGet(topoMap, newPos))
                {
                    // Console.WriteLine($"({newPos.col}, {newPos.row}), ({TopoMapGet(topoMap, current) + 1}| {TopoMapGet(topoMap, newPos)})");
                    var newList = walk(topoMap, newPos);
                    result.AddRange(newList);
                }
            }

            return result;
        }

    }
}
