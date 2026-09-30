package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gifs := []string{
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatPat.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatPawPsychadelia.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatPopsBaloon.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatStuckTongue.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatsAndMops.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/ChangingShiftsInTheBox.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/Corgi%20Run%20Stop.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/DeterminedToiletCat.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/Dog%20Car%20Mirror%20Freak.gif",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		count := 1
		if value := r.URL.Query().Get("count"); value != "" {
			parsed, err := strconv.Atoi(value)
			if err == nil && parsed >= 1 && parsed <= 3 {
				count = parsed
			}
		}

		// Choose unique random GIFs for the requested count.
		indices := rand.Perm(len(gifs))
		selected := make([]string, 0, count)
		for i := 0; i < count; i++ {
			selected = append(selected, gifs[indices[i]])
		}

		selectedOption := func(n int) string {
			if count == n {
				return " selected"
			}
			return ""
		}

		cardsHTML := ""
		for _, gifURL := range selected {
			cardsHTML += fmt.Sprintf(`<div class="card"><img src="%s" alt="Cute animal GIF"></div>`, gifURL)
		}

		var html strings.Builder
		html.WriteString(`<!DOCTYPE html>
			<html>
			<head>
				<meta charset="utf-8">
				<title>Random GIF Activity</title>
				<style>
					body {
						margin: 0;
						background: linear-gradient(135deg, #dff7ff 0%, #f7f7ff 100%);
						font-family: Arial, sans-serif;
						padding: 2rem;
						color: #1f3b5b;
					}
					.container {
						max-width: 1100px;
						margin: 0 auto;
					}
					h1 {
						margin-top: 0;
						text-align: center;
						color: #1f3b5b;
					}
					.controls {
						display: flex;
						justify-content: center;
						align-items: center;
						gap: 0.75rem;
						margin: 1.5rem 0 2rem;
						flex-wrap: wrap;
					}
					label {
						font-weight: bold;
					}
					select {
						padding: 0.6rem 0.9rem;
						font-size: 1rem;
						border: 1px solid #8ecae6;
						border-radius: 8px;
						background: white;
					}
					button {
						padding: 0.6rem 1rem;
						font-size: 1rem;
						background: #1f3b5b;
						color: white;
						border: none;
						border-radius: 8px;
						cursor: pointer;
					}
					.grid {
						display: grid;
						grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
						gap: 1.25rem;
						align-items: center;
					}
					.card {
						background: rgba(255, 255, 255, 0.5);
						border-radius: 16px;
						padding: 0.75rem;
						box-shadow: 0 10px 20px rgba(31, 59, 91, 0.08);
						display: flex;
						justify-content: center;
						align-items: center;
						min-height: 220px;
					}
					img {
						display: block;
						width: 100%;
						height: 220px;
						object-fit: cover;
						border-radius: 12px;
						box-shadow: 0 8px 18px rgba(0, 0, 0, 0.12);
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Hello, world!</h1>
					<form class="controls" method="get" action="/">
						<label for="count">Number of GIFs:</label>
						<select id="count" name="count">
							<option value="1"`)
		html.WriteString(selectedOption(1))
		html.WriteString(`>1</option>
							<option value="2"`)
		html.WriteString(selectedOption(2))
		html.WriteString(`>2</option>
							<option value="3"`)
		html.WriteString(selectedOption(3))
		html.WriteString(`>3</option>
						</select>
						<button type="submit">Show GIFs</button>
					</form>
					<div class="grid">`)
		html.WriteString(cardsHTML)
		html.WriteString(`</div>
				</div>
			</body>
			</html>`)

		if _, err := w.Write([]byte(html.String())); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
