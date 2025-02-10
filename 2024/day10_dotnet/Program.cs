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
                    var newP = new Point { row = i, col = j };
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

        int counter = 0;
        string horientation = "down";

        foreach (var start in startPoints)
        {
            Console.WriteLine(walk(topoMap, start, horientation, counter, seen));
        }
    }

    static bool walk(List<List<char>> todoMap, Point start, string horientation, int counter, List<List<bool>> seen)
    {
        Console.WriteLine("Entra");
        return true;
    }
}
