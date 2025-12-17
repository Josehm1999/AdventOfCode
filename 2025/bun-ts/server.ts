const server = Bun.serve({
  async fetch(req) {
    const path = new URL(req.url).pathname;

    // Serve your HTML
    if (path === "/") {
      return new Response(Bun.file("./visualize.html"), {
        headers: { "Content-Type": "text/html" },
      });
    }

    // Serve your JS
    if (path === "/draw.js") {
      return new Response(Bun.file("./draw.js"), {
        headers: { "Content-Type": "application/javascript" },
      });
    }

    // Serve your data file
    if (path === "/test_day9.txt") {
      return new Response(Bun.file("./test_day9.txt"), {
        headers: { "Content-Type": "text/plain" },
      });
    }

    return new Response("Not found", { status: 404 });
  },
});

console.log(`Listening on ${server.url}`);
