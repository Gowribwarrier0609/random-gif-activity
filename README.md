# Random GIF activity

Build a tiny Go web app that shows a GIF. You'll work in pairs and take turns
making changes. The point is to see how a **server** handles a browser's HTTP
request, reads information from that request, and sends back HTML.

## Get started

1. Find a partner! Meet somebody new. Make a [human] friend.
2. Fork this repo.
4. Open this repository in a GitHub Codespace (**Code → Codespaces → Create
   codespace on main**) or clone it to a computer with Go installed.
5. In the terminal, run `go run .`.
6. Open `http://localhost:8080`. In a Codespace, use the **Ports** tab to open
   port 8080 in your browser. You should see `Hello, world!`.
7. Open `main.go`. The function passed to `http.HandleFunc` handles requests
   for `/`. The `w` value is the response sent to the browser; `r` is the
   incoming request. Change the greeting, save the file, stop the server with
   Ctrl-C, and run `go run .` again to see your change.
8. Now you have the basic app working! Ok, where to go next?
9. File issues based on the user stories below. (You can likely use the built-in co-pilot AI; use your own AI; or do this manually. If using your own AI, you can use the GitHub MCP, the GitHub `gh` CLI, or use their API. So many options! You *might* be able to sign into Yale SOM's Claude Code account from the CodeSpaces terminal.)
10. Complete each issue. You can work on branches or commit directly to `main`. Use branches or not. It's likely not a big deal for this small project. Make sure the "acceptance criteria" for your issues include tests.
11. 

The [GIF list](gifs.txt) links to 146 animal GIFs from the
[`adorbs` collection](https://github.com/snipe/animated-gifs/tree/master/adorbs).
The images are hosted in course storage, so you can use them without an API key.
Open a URL from the list in your browser to see what it shows.

## User stories

Work through these in order. We'll pause after each one to compare approaches.

1. As a visitor, I see **Hello, world!** at `/`. The starter already does this.
2. As a visitor, I see an HTML page with a heading and a colored background.
   Before writing the HTML response, set its content type with
   `w.Header().Set("Content-Type", "text/html; charset=utf-8")`.
3. As a visitor, I see a GIF on the page. Copy a URL from `gifs.txt` into an
   HTML image tag, such as `<img src="GIF_URL" alt="A cute animal">`.
4. As a visitor, I get a different GIF when I reload the page. Put at least
   two GIF URLs in a Go slice and choose one in the request handler.
   Browse `gifs.txt` and pick your favorites. Or choose them all!
5. As a visitor, I can enter `cat` or `puppy` in a search box and submit it.
   The server reads the query from `r.URL.Query().Get("q")` and chooses from
   the GIFs you assigned to that word. An HTML form with `method="get"` and
   `action="/"` will send the query as `/?q=cat`. For an unknown or empty
   query, show a random GIF or a helpful message.

The server handles the form and chooses a URL; the browser then asks course
storage for that GIF. No API key or image proxy is needed.

## Check your work

- Refreshing `/` changes the GIF at least some of the time.
- Submitting `cat` produces a URL containing `?q=cat` and a cat GIF.
- Submitting an unknown word does not crash the server.
- The page still works after stopping and restarting `go run .`.

If you have time, add another way to filter the GIFs or make the page look
nicer. This is an in-class exercise; you do not need to deploy or submit it.

## GIF credits

The GIFs are from [Snipe's `adorbs` collection](https://github.com/snipe/animated-gifs/tree/d5ff840d028c2438497e7a7709d6bb9d5f7c6d68/adorbs)
at commit `d5ff840d028c2438497e7a7709d6bb9d5f7c6d68`.
The [source README](https://github.com/snipe/animated-gifs/blob/master/README.md)
credits the respective image copyright holders; it does not provide a license.
Course copies expire after December 30, 2026.
