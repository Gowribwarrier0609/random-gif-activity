# Random GIF activity

Build a tiny Go web app that shows a GIF. You'll work in pairs and take turns
making changes. The point is to see how a **server** handles a browser's HTTP
request, reads information from that request, and sends back HTML.

## Get started

1. Open this repository in a GitHub Codespace (**Code → Codespaces → Create
   codespace on main**) or clone it to a computer with Go installed.
2. In the terminal, run `go run .`.
3. Open `http://localhost:8080`. In a Codespace, use the **Ports** tab to open
   port 8080 in your browser. You should see `Hello, world!`.
4. Open `main.go`. The function passed to `http.HandleFunc` handles requests
   for `/`. The `w` value is the response sent to the browser; `r` is the
   incoming request. Change the greeting, save the file, stop the server with
   Ctrl-C, and run `go run .` again to see your change.

The repository includes eleven cat, dog, and parrot GIFs in `gifs/`. The
starter already serves that directory: open `/gifs/black-cat-roll.gif` in your
browser to check. The GIFs are local files, so the app does not need Giphy or
another image service.

## User stories

Work through these in order. We'll pause after each one to compare approaches.

1. As a visitor, I see **Hello, world!** at `/`. The starter already does this.
2. As a visitor, I see an HTML page with a heading and a colored background.
   Before writing the HTML response, set its content type with
   `w.Header().Set("Content-Type", "text/html; charset=utf-8")`.
3. As a visitor, I see a GIF on the page. An HTML image tag looks like
   `<img src="/gifs/black-cat-roll.gif" alt="A black cat rolling">`.
4. As a visitor, I get a different GIF when I reload the page. Put at least
   two local GIF paths in a Go slice and choose one in the request handler.
   Browse the `gifs/` directory and pick your favorites.
5. As a visitor, I can enter `cat`, `dog`, or `parrot` in a search box and
   submit it. The server reads the query from `r.URL.Query().Get("q")` and
   chooses a GIF from that category. An HTML form with `method="get"` and
   `action="/"` will send the query as `/?q=dog`. For an unknown or empty
   query, show a random GIF or a helpful message.

The server handles the form and chooses a local image; the browser then asks
the same server for that GIF. No API key, external proxy, or image CDN is needed.

## Check your work

- Refreshing `/` changes the GIF at least some of the time.
- Submitting `dog` produces a URL containing `?q=dog` and a dog GIF.
- Submitting an unknown word does not crash the server.
- The page still works after stopping and restarting `go run .`.

If you have time, add another way to filter the GIFs or make the page look
nicer. This is an in-class exercise; you do not need to deploy or submit it.

## GIF credits

The GIFs come from [Mochi](https://github.com/koustavdatascience/mochi),
[SalaryCat](https://github.com/Einswen/SalaryCat), and
[Plant Pets](https://opengameart.org/content/plant-pets). See
[`gifs/CREDITS.md`](gifs/CREDITS.md) for the source and license of each file.
The Go starter and activity instructions are separate from those image assets.
